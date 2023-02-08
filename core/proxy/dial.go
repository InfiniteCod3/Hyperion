package proxy

import "net"

type Dial struct {
	Dial func(string, string) (net.Conn, error)
}
