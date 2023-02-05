package methods

import "Hyperion/core/method"

type Join struct{}

func (join Join) Name() string {
	return "Join"
}

func (join Join) Description() string {
	return "Join method implementation"
}

func (join Join) Start(info method.AttackInfo) {
	// implementation for starting the join method
}

func (join Join) Stop() {
	// implementation for stopping the join method
}
