package main

import (
	"os"
	"time"

	"github.com/lunarxlark/learn-go-with-tests/mocking"
)

func main() {
	//sleeper := mocking.DefaultSleeper{}
	sleeper := *mocking.ConfigurableSleeper{1 * time.Second, time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}