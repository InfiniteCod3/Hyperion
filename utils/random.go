package utils

import (
	"math/rand"
	"time"
)

var (
	validNameRunes    = []rune("aAbBcCdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ0123456789_")
	validNameRunesLen = len(validNameRunes)
	randCh            = make(chan int, 1024)
	src               = rand.NewSource(time.Now().UnixNano())
)

func Init() {
	go func() {
		batch := make([]int, 1024)
		for {
			for i := range batch {
				batch[i] = int(src.Int63() % int64(validNameRunesLen))
			}
			for _, num := range batch {
				randCh <- num
			}
		}
	}()
}

func RandomName(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = validNameRunes[<-randCh]
	}
	return string(b)
}
