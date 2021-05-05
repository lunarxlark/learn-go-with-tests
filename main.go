package main

import (
	"os"

	"github.com/lunarxlark/learn-go-with-tests/mocking"
)

func main() {
	sleeper := mocking.DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)
}