package utils

import (
	"math/rand"
	"strings"
	"time"
)

var valid_name_runes = []rune("aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789_")
var sb = strings.Builder{}

func seed() {
	rand.Seed(time.Now().UnixNano())
}

// func RandomUTF16String(length int) (str string) {
// 	seed()
// }

// func RandomUTF8String(length int) (str string) {
// 	seed()
// }

func RandomName(length int) (str string) {
	seed()
	for i := 0; i < length; i++ {
		sb.WriteRune(valid_name_runes[rand.Intn(len(valid_name_runes))])
	}
	str = sb.String()
	sb.Reset()
	return
}

// func RandomNumberString(lowerBound, upperBound int) (str string) {
// 	seed()
// }
