package methods

import (
	"Hyperion/core"
	"Hyperion/core/proxy"
)

type Ping struct {
	Info         *core.AttackInfo
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
