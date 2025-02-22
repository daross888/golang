package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}

func Countdown(writer io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		s.Sleep()
	}

	fmt.Fprint(writer, finalWord)
}
