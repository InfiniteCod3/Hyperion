package method

import (
	"strings"
	"time"
)

type AttackInfo struct {
	Ip       string
	Port     int
	Protocol int
	Duration time.Duration
	Delay    time.Duration
	PerDelay time.Duration
}

type Method interface {
	Name() string
	Description() string
	Start(info AttackInfo)
	Stop()
}

var (
	methods = make(map[string]Method)
)

func RegisterMethod(method Method) {
	methods[strings.ToLower(method.Name())] = method
}

func GetMethod(name string) (method Method, ok bool) {
	method, ok = methods[strings.ToLower(name)]
	return
}
