package proxy

import (
	"strconv"
)

type ProxyProtocol string

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

func (manager *ProxyManager) GetNext() (proxy *Proxy) {
	proxy = manager.proxies[manager.atIndex]
	manager.atIndex = (manager.atIndex + 1) % len(manager.proxies)
	return
}
