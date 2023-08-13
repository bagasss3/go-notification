package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func ParseTimeDuration(t string, defaultt time.Duration) time.Duration {
	timeDurr, err := time.ParseDuration(t)
	if err != nil {
		return defaultt
	}
	return timeDurr
}

func GenerateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano()+int64(rand.Intn(10000)))
}
