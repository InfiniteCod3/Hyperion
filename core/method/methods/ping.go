package methods

import (
	"Hyperion/core/method/methods"
	"Hyperion/core/method/methods"
	"Hyperion/core/proxy"
)

type Ping struct {
	Info         *methods.AttackInfo
	ProxyManager *proxy.ProxyManager
}

func (ping Ping) Name() string {
	return "Ping"
}

func (ping Ping) Description() string {
	return "Flood server with pings"
}

func (ping Ping) Start() {

}

func (ping Ping) Stop() {
	// implementation for stopping the Ping method
}
