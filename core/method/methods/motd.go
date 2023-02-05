package methods

import (
	"Hyperion/core/method"
)

type MOTD struct{}

func (Method MOTD) Name() string {
	return "MOTD"
}

func (j MOTD) Description() string {
	return "Motd method implementation"
}

func (motd MOTD) Start(info method.AttackInfo) {
	// implementation for starting the Motd method
}

func (motd MOTD) Stop() {
	// implementation for stopping the Motd method
}
