package method

import (
	"Hyperion/core"
	"strings"
)

type Method interface {
	Name() string
	Description() string
	Start(info core.AttackInfo)
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
