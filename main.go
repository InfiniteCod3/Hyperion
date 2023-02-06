package main

import (
	"Hyperion/core"
	"Hyperion/core/method"
	"Hyperion/core/method/methods"
	"flag"
	"fmt"
	"os"
)

var (
	ip       = flag.String("ip", "0.0.0.0", "sets ip")
	port     = flag.Int("port", 25565, "sets port")
	protocol = flag.Int("protcol", 761, "sets version protocol")
)

func main() {
	flag.Parse()
	registerMethod()
	method, ok := method.GetMethod("join")
	if !ok {
		fmt.Println("ERROR GETTING METHOD")
		os.Exit(1)
	}
	method.Start(core.AttackInfo{
		Ip:       *ip,
		Port:     *port,
		Protocol: *protocol,
	})
}

func registerMethod() {
	method.RegisterMethod(methods.Join{})
	method.RegisterMethod(methods.Ping{})
	method.RegisterMethod(methods.MOTD{})
}
