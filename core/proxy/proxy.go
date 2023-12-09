package proxy

import (
	"net"
	"strconv"
)

type ProxyProtocol int

const (
	SOCKS4  ProxyProtocol = 0
	SOCKS4A ProxyProtocol = 1
	SOCKS5  ProxyProtocol = 2
)

type Proxy struct {
	Ip       string
	Port     string
	Protocol ProxyProtocol
}

func (proxy *Proxy) GetString() (key string) {
	return strconv.Itoa(int(proxy.Protocol)) + "://" + net.JoinHostPort(proxy.Ip, proxy.Port)
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

func (manager *ProxyManager) GetNext() *Proxy {
	manager.atIndex = manager.atIndex + 1
	if manager.atIndex >= len(manager.proxies) {
		manager.atIndex = 0
	}
	return manager.proxies[manager.atIndex]
}
