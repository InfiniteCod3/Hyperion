package methods

import (
	"Hyperion/core/method/methods"
	"Hyperion/core/proxy"
	"Hyperion/mc"
	"Hyperion/mc/mcutils"
)

type MOTD struct {
	Info         *method.AttackInfo
	ProxyManager *proxy.ProxyManager
}

func (motd MOTD) Name() string {
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
