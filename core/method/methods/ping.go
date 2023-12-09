package methods

import (
	"Hyperion/core"
	"Hyperion/core/proxy"
)

type Ping struct {
	Info         *core.AttackInfo
	ProxyManager *proxy.ProxyManager
}

func (ping Ping) Name() string {
	return "Ping"
}

func (ping Ping) Description() string {
	return "Flood server with pings"
}

func (ping Ping) Start() {
	// Create a worker pool
	workerPool := make(chan struct{}, 10) // 10 is the number of workers
	stop := make(chan struct{})

	// Start the workers
	for i := 0; i < 10; i++ {
		go func() {
			for {
				select {
				case <-stop:
					return
				case workerPool <- struct{}{}:
					// Get the next proxy
					proxy := ping.ProxyManager.GetNext()
					// Do the work
					// ...
					// When done, release the worker
					<-workerPool
				}
			}
		}()
	}
}

func (ping Ping) Stop() {
	// implementation for stopping the Ping method
}
}
