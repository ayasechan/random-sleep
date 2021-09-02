package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	maxSleep   = flag.Duration("max", 0, "maximum delay")
	minSleep   = flag.Duration("min", 0, "minimum delay")
	isRaiseErr = flag.Int("raise-error", 0, "probability of abnormal exit. between 0-100")
	noPrint    = flag.Bool("no-print", false, "prohibit the printing delay time")
)

func main() {
	flag.Parse()
	if *maxSleep == 0 {
		flag.Usage()
		os.Exit(1)
	}
	exitCode := 0
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(100+1) < *isRaiseErr {
		exitCode = 1
	}
	duration := GetRandomDuration(*maxSleep, *minSleep)
	if !*noPrint {
		fmt.Println(duration)
	}
	time.Sleep(duration)
	os.Exit(exitCode)
}

func GetRandomDuration(max, min time.Duration) time.Duration {
	if max == min {
		return max
	}
	if max < min {
		max, min = min, max
	}
	maxSeconds := int(max.Seconds()) + 1
	minSeconds := int(min.Seconds())
	rand.Seed(time.Now().UnixNano())
	randomSeconds := rand.Intn(maxSeconds-minSeconds) + minSeconds
	return time.Second * time.Duration(randomSeconds)
}
