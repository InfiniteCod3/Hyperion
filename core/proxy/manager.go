package proxy

import (
	"net"
	"strconv"

	"h12.io/socks"
)

type ProxyProtocol string

type Dial func(string, string) (net.Conn, error)

const (
	HTTP   ProxyProtocol = "http"
	SOCKS4 ProxyProtocol = "socks4"
	SOCKS5 ProxyProtocol = "socks5"
)

type Proxy struct {
	Ip       string
	Port     uint16
	Protocol ProxyProtocol
}

func (proxy *Proxy) GetString() (key string) {
	return string(proxy.Protocol) + "://" + proxy.Ip + ":" + strconv.Itoa(int(proxy.Port))
}

type ProxyManager struct {
	proxies []*Proxy
}

func (manager *ProxyManager) Add(proxy *Proxy) {
	manager.proxies = append(manager.proxies, proxy)
}

func (manager *ProxyManager) Remove(proxy *Proxy) {
	for i, p := range manager.proxies {
		if p.GetString() == proxy.GetString() {
			manager.proxies = append(manager.proxies[:i], manager.proxies[i+1:]...)
			return
		}
	}
}

type DialPool struct {
	dials   []*Dial
	atIndex int
}

func (manager *DialPool) GetNext() (dial *Dial) {
	dial = manager.dials[manager.atIndex]
	manager.atIndex = (manager.atIndex + 1) % len(manager.dials)
	return
}

func (manager *DialPool) AddFromProxyManager(proxyManager *ProxyManager) {
	for _, proxy := range proxyManager.proxies {
		dial := Dial(socks.Dial(proxy.GetString()))
		manager.Add(&dial)
	}
}

func (manager *DialPool) Add(dial *Dial) {
	manager.dials = append(manager.dials, dial)
}

func (manager *DialPool) Remove(dial *Dial) {
	for i, d := range manager.dials {
		if d == dial {
			manager.dials = append(manager.dials[:i], manager.dials[i+1:]...)
			return
		}
	}
}
