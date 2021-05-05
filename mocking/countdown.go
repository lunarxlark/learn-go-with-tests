package mocking

import (
	"fmt"
	"io"
	"time"
)

const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(writer io.Writer) {
	for i := 3; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintf(writer, "%d\n", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprintf(writer, finalWord)
}
