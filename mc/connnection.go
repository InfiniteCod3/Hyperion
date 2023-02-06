package mc

import (
	pk "Hyperion/mc/packet"
	"context"
	"crypto/cipher"
	"errors"
	"io"
	"net"
	"strconv"
	"time"
)

const DefaultPort = 25565

type Listener struct{ net.Listener }

func ListenMC(addr string) (*Listener, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Listener{l}, nil
}

func (l Listener) Accept() (Connection, error) {
	conn, err := l.Listener.Accept()
	return Connection{
		Socket:    conn,
		Reader:    conn,
		Writer:    conn,
		threshold: -1,
	}, err
}

type Connection struct {
	Socket net.Conn
	io.Reader
	io.Writer

	threshold int
}

var DefaultDialer = Dialer{}

func DialMC(addr string) (*Connection, error) {
	return DefaultDialer.DialMCContext(context.Background(), addr)
}

func DialMCTimeout(addr string, timeout time.Duration) (*Connection, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return DefaultDialer.DialMCContext(ctx, addr)
}

type Dialer net.Dialer

func (d *Dialer) resolver() *net.Resolver {
	if d != nil && d.Resolver != nil {
		return d.Resolver
	}
	return net.DefaultResolver
}

func (d *Dialer) DialMCContext(ctx context.Context, addr string) (*Connection, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		var addrErr *net.AddrError
		const missingPort = "missing port in address"
		if errors.As(err, &addrErr) && addrErr.Err == missingPort {
			host, port, err = addr, "", nil
		} else {
			return nil, err
		}
	}
	var ras []string
	if port == "" {
		_, srvRecords, err := d.resolver().LookupSRV(ctx, "minecraft", "tcp", host)
		if err == nil {
			for _, record := range srvRecords {
				addr := net.JoinHostPort(record.Target, strconv.Itoa(int(record.Port)))
				ras = append(ras, addr)
			}
		}
		addr = net.JoinHostPort(addr, strconv.Itoa(DefaultPort))
	}
	ras = append(ras, addr)

	var firstErr error
	for i, addr := range ras {
		select {
		case <-ctx.Done():
			return nil, context.Canceled
		default:
		}
		dialCtx := ctx
		if deadline, hasDeadline := ctx.Deadline(); hasDeadline {
			partialDeadline, err := partialDeadline(time.Now(), deadline, len(ras)-i)
			if err != nil {
				if firstErr == nil {
					firstErr = context.DeadlineExceeded
				}
				break
			}
			if partialDeadline.Before(deadline) {
				var cancel context.CancelFunc
				dialCtx, cancel = context.WithDeadline(ctx, partialDeadline)
				defer cancel()
			}
		}
		conn, err := (*net.Dialer)(d).DialContext(dialCtx, "tcp", addr)
		if err != nil {
			if firstErr == nil {
				firstErr = err
			}
			continue
		}
		return WrapConn(conn), nil
	}
	return nil, firstErr
}

func partialDeadline(now, deadline time.Time, addrsRemaining int) (time.Time, error) {
	if deadline.IsZero() {
		return deadline, nil
	}
	timeRemaining := deadline.Sub(now)
	if timeRemaining <= 0 {
		return time.Time{}, context.DeadlineExceeded
	}
	timeout := timeRemaining / time.Duration(addrsRemaining)
	const saneMinimum = 2 * time.Second
	if timeout < saneMinimum {
		if timeRemaining < saneMinimum {
			timeout = timeRemaining
		} else {
			timeout = saneMinimum
		}
	}
	return now.Add(timeout), nil
}

func WrapConn(conn net.Conn) *Connection {
	return &Connection{
		Socket:    conn,
		Reader:    conn,
		Writer:    conn,
		threshold: -1,
	}
}

func (c *Connection) Close() error { return c.Socket.Close() }

func (c *Connection) ReadPacket(p *pk.Packet) error {
	return p.UnPack(c.Reader, c.threshold)
}

func (c *Connection) WritePacket(p pk.Packet) error {
	return p.Pack(c.Writer, c.threshold)
}

func (c *Connection) SetCipher(ecoStream, decoStream cipher.Stream) {
	c.Reader = cipher.StreamReader{ // Set receiver for AES
		S: decoStream,
		R: c.Socket,
	}
	c.Writer = cipher.StreamWriter{
		S: ecoStream,
		W: c.Socket,
	}
}

func (c *Connection) SetThreshold(t int) {
	c.threshold = t
}
