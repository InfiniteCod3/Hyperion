package methods

import (
	"Hyperion/core"
	"Hyperion/core/proxy"
)

type MOTD struct {
	Info         *core.AttackInfo
	ProxyManager *proxy.ProxyManager
}

func (Method MOTD) Name() string {
	return "MOTD"
}

func (motd MOTD) Description() string {
	return "Joins server and then flood request motd"
}

func (motd MOTD) Start() {

}

func (motd MOTD) Stop() {
	// implementation for stopping the Motd method
}
