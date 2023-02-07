package methods

import (
	"Hyperion/core"
	"Hyperion/core/proxy"
)

type MOTD struct{}

func (Method MOTD) Name() string {
	return "MOTD"
}

func (j MOTD) Description() string {
	return "Joins server and then flood request motd"
}

func (motd MOTD) Start(info *core.AttackInfo, dialPool *proxy.DialPool) {

}

func (motd MOTD) Stop() {
	// implementation for stopping the Motd method
}
