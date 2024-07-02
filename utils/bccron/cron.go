package bccron

import (
	"math"
	"time"
)

func Next(cron string) int {
	now := time.Now()
	diff := MustParse(cron).Next(now).Sub(now).Seconds()
	if diff < 0 {
		return 0
	}
	return int(math.Ceil(diff))
}
