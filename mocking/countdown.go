package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{duration: 1 * time.Second, sleep: time.Sleep})
}

// Countdown counts down to Go!
func Countdown(w io.Writer, s Sleeper) {
	for i := 3; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprint(w, "Go!")
}

// Sleeper is something that sleeps
type Sleeper interface {
	Sleep()
}

// CountdownOperationsSpy spies on all ops that happen
type CountdownOperationsSpy struct {
	Calls []string
}

// Sleep adds sleep to list of ops
func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
	return
}

// Write adds write to list of ops
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const write = "write"
const sleep = "sleep"

// ConfigurableSleeper takes a sleep func
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// SpyTime is for storing a time we were asked to sleep for
type SpyTime struct {
	durationSlept time.Duration
}

// Sleep stores duration "slept" by a spyTime
func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// Sleep sleeps by calling the sleep() func stored in the struct
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)

}
