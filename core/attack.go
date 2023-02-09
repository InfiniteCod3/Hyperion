package core

import (
	"time"
)

type AttackInfo struct {
	Ip                   string
	Port                 string
	Protocol             int
	Duration             time.Duration
	ConnPerProxyPerDelay int
	Delay                time.Duration
}
