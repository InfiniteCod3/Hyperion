package methods

import (
	"Hyperion/core/method"
)

type Ping struct{}

func (ping Ping) Name() string {
	return "Ping"
}

func (ping Ping) Description() string {
	return "Ping method implementation"
}

func (ping Ping) Start(info method.AttackInfo) {
	// implementation for starting the Ping method
}

func (ping Ping) Stop() {
	// implementation for stopping the Ping method
}
