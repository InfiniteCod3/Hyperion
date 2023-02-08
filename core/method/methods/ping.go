package methods

import (
	"Hyperion/core"
	"Hyperion/core/proxy"
)

type Ping struct {
	Info         *core.AttackInfo
	ProxyManager *proxy.ProxyManager
	isRunning    bool
}

func (ping Ping) Name() string {
	return "Ping"
}

func (ping Ping) Description() string {
	return "Flood server with pings"
}

func (ping Ping) IsRunning() bool {
	return ping.isRunning
}

func (ping Ping) Start() {

}

func (ping Ping) Stop() {
	// implementation for stopping the Ping method
}
