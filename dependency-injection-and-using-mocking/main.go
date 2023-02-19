package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper allows you to put delays.
type Sleeper interface {
	Sleep()
}

// DefaultSleeper is an implementation of Sleeper with a predefined delay.
type DefaultSleeper struct {
	// optional - adding duration to sleep that will be used in Sleep() implementation
	duration time.Duration
	sleep    func(time.Duration) // a func that will sleep for specified duration. time.Sleep provides this capability
}

// Sleep will pause execution for the defined Duration.
func (d *DefaultSleeper) Sleep() {
	//time.Sleep(1 * time.Second)
	d.sleep(d.duration)
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints a countdown from 3 to out with a delay between count provided by Sleeper.
func Countdown(out io.Writer, sleeper Sleeper) {
	// injecting dependency io.Writer (passing value io.Writer) to also allow testing
	// instead of printing to output
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
