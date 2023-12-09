package method

import (
	"strings"
)

type Method interface {
	Name() string
	Description() string
	Start()
	Stop()
}

var (
	methods = make(map[string]Method)
)

func RegisterMethod(method Method) error {
	if _, ok := methods[strings.ToLower(method.Name())]; ok {
		return fmt.Errorf("method already exists")
	}
	methods[strings.ToLower(method.Name())] = method
	return nil
}

func GetMethod(name string) (method Method, ok bool) {
	method, ok = methods[strings.ToLower(name)]
	return
}
