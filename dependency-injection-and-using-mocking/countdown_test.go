package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		// can't use bytes.Buffer{} as io.Writer has a Write method that is called by print
		// which takes a pointer receiver
		buffer := &bytes.Buffer{}
		// can't use SpyCountDownOpsSleeper{} as we have to do s.Calls operation so pointer needed
		// spySleeper := &SpySleeper{}

		// we can use &SpyCountDownOpsSleeper{}, as the only thing we need is .Sleep() method and &SpyCountDownOpsSleeper{} has that
		Countdown(buffer, &SpyCountDownOpsSleeper{})

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			// t testing.T instance methods like Error() for testing purpose
			// Use Errorf() as Error() doesn't do formatting things like %q.
			// Instead, it uses the default format of its arguments, and adds spaces between them
			t.Errorf("got %q want %q", got, want)
		}

	})

	t.Run("sleep in-betw prints", func(t *testing.T) {
		// create SpySleeper to test sleep functionality that has the same methods (by implementing those) as DefaultSleeper
		// and implement behavior of those methods in a manner that it allows to test sleep in-betw prints functionality
		spySleeper := &SpyCountDownOpsSleeper{}

		// we can use &SpyCountDownOpsSleeper{}, as the only thing we need is .Sleep() and .Write() methods
		// and &SpyCountDownOpsSleeper{} has them
		Countdown(spySleeper, spySleeper)

		want := []string{
			write, sleep, write, sleep, write, sleep, write,
		}

		if !reflect.DeepEqual(want, spySleeper.Calls) {
			t.Errorf("wanted calls %q got %q", want, spySleeper.Calls)
		}

	})

}

func TestConfigurableSleepDurationFunctionality(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	// create sleeper to pass in CountDown using a spy sleep method, to test the sleep duration functionality
	// we can use &SpyTime as the only thing we need is Sleep func to create DefaultSleeper and call .Sleep() with it
	// and &SpyTime{} has that
	sleeper := DefaultSleeper{sleepTime, spyTime.Sleep}

	// call actual Sleep method which must invoke spyTime.Sleep(sleepTime) with sleepTime.
	// if correct behavior, this will set the spyTime.durationSlept as sleepTime as defined
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}

}

type SpyCountDownOpsSleeper struct {
	Calls []string // initialised to empty array if not passes while declaring SpyCountDownOpsSleeper
}

func (s *SpyCountDownOpsSleeper) Write(p []byte) (n int, err error) {
	// implement Write() method of io.Writer used by print
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyCountDownOpsSleeper) Sleep() {
	// implement Sleep method of DefaultSleeper{} present in main.go
	s.Calls = append(s.Calls, sleep)
}

const write = "write"
const sleep = "sleep"

type SpyTime struct {
	durationSlept time.Duration // data members of class SpyTime
}

func (s *SpyTime) Sleep(duration time.Duration) { // methods of class SpyTime
	s.durationSlept = duration
}
