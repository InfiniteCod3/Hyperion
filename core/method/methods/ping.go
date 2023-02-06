package methods

import (
	"Hyperion/core"
)

type Ping struct{}

func (ping Ping) Name() string {
	return "Ping"
}

func (ping Ping) Description() string {
	return "Flood server with pings"
}

func (ping Ping) Start(info core.AttackInfo) {

}

func (ping Ping) Stop() {
	// implementation for stopping the Ping method
}
