package main

import (
	"fmt"
	"time"
)

func main() {
	chronosTime := getTime()
	fmt.Println(chronosTime)
}

func getTime() (chronosTime string) {
	now := time.Now().UTC()
	midnight := now.Truncate(24 * time.Hour)
	nanosecondsSinceMidnight := time.Since(midnight).Nanoseconds()
	proportionOfDay := float64(nanosecondsSinceMidnight) / float64(24*60*60*1000000000)
	chronosTime = fmt.Sprintf("%04X", int64(proportionOfDay*65536))
	return chronosTime
}
