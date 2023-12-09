package mc

import (
	"Hyperion/core/proxy"
	pk "Hyperion/mc/packet"
	"crypto/cipher"
	"io"
	"net"

	"h12.io/socks"
)

const DefaultPort = 25565

type Connection struct {
	Socket net.Conn
	Reader io.Reader
	Writer io.Writer

	threshold int
}

func DialMC(ip *string, port *string, proxy *proxy.Proxy) (connection *Connection, err error) {
	conn, perr := socks.DialSocksProxy(int(proxy.Protocol), net.JoinHostPort(proxy.Ip, proxy.Port))("tcp", net.JoinHostPort(*ip, *port))
	err = perr
	connection = WrapConn(conn)
	return
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
