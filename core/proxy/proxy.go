package proxy

import (
	"net"
)

type ProxyProtocol string

const (
	HTTP   ProxyProtocol = "http"
	SOCKS4 ProxyProtocol = "socks4"
	SOCKS5 ProxyProtocol = "socks5"
)

type Proxy struct {
	Ip       string
	Port     string
	Protocol ProxyProtocol
}

func (proxy *Proxy) GetString() (key string) {
	return string(proxy.Protocol) + "://" + net.JoinHostPort(proxy.Ip, proxy.Port)
}

type ProxyManager struct {
	proxies []*Proxy
	atIndex int
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

func (manager *ProxyManager) Length() (length int) {
	length = len(manager.proxies)
	return
}

func (manager *ProxyManager) GetNext() (proxy *Proxy) {
	if manager.atIndex >= len(manager.proxies) {
		manager.atIndex = 0
	} else {
		manager.atIndex = manager.atIndex + 1
	}
	proxy = manager.proxies[manager.atIndex]
	return
}
