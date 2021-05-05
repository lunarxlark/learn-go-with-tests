package main

import (
	"os"

	"github.com/lunarxlark/learn-go-with-tests/mocking"
)

func main() {
	mocking.Countdown(os.Stdout)
}