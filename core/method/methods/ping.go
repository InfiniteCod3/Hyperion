package methods

import (
	"Hyperion/core"
	"Hyperion/core/proxy"
	"sync"
)

type Ping struct {
	connPool sync.Pool
}

func (ping Ping) Name() string {
	return "Ping"
}

func (ping Ping) Description() string {
	return "Flood server with pings"
}

func (ping Ping) Start(info *core.AttackInfo, dialPool *proxy.DialPool) {

}

func (ping Ping) Stop() {
	// implementation for stopping the Ping method
}
