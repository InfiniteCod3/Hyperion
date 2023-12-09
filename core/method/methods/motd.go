package methods

import (
	"Hyperion/core"
	"Hyperion/core/proxy"
)

type MOTD struct {
	Info         *core.AttackInfo
	ProxyManager *proxy.ProxyManager
}

func (Method MOTD) Name() string {
	return "MOTD"
}

func (motd MOTD) Description() string {
	return "Joins server and then flood request motd"
}

func (motd MOTD) Start() {
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
					proxy := motd.ProxyManager.GetNext()
					// Do the work
					// ...
					// When done, release the worker
					<-workerPool
				}
			}
		}()
	}
}

func (motd MOTD) Stop() {
	// implementation for stopping the Motd method
}
	close(stop)
}
