package core

import "time"

type AttackInfo struct {
	Ip       string
	Port     int
	Protocol int
	Duration time.Duration
	Delay    time.Duration
	PerDelay int
}
