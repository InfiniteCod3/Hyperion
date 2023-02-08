package methods

import (
	"Hyperion/core"
	"Hyperion/core/proxy"
	"sync"
)

type MOTD struct {
	connPool sync.Pool
}

func (Method MOTD) Name() string {
	return "MOTD"
}

func (j MOTD) Description() string {
	return "Joins server and then flood request motd"
}

func (motd MOTD) Start(info *core.AttackInfo, proxyManager *proxy.ProxyManager) {

}

func (motd MOTD) Stop() {
	// implementation for stopping the Motd method
}
