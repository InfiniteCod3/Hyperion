package utils

import (
	"math/rand"
	"time"
)

var validNameRunes = []rune("aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789_")

func Init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomName(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = validNameRunes[rand.Intn(len(validNameRunes))]
	}
	return string(b)
}
