package core

import (
	"time"
)

type AttackInfo struct {
	Ip           string
	Port         string
	Protocol     int
	Duration     time.Duration
	ConnPerProxy int
	Delay        time.Duration
	Loops        int
}
