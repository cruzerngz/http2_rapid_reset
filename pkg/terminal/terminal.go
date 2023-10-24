// For manipulating terminal stuffs

package terminal

import (
	"fmt"
	"time"
)

const UINT_MAX_TERMINATOR uint = ^uint(0)
const PREVIOUS_LINE string = "\033[F"

// Events tally counter
type Events struct {
	header    string
	numEvents uint64
	freq      float32
	channel   chan uint
}

// Create a new event tracker along with its attached channel
func NewEvent(title string) (*Events, chan uint) {

	var ch chan uint = make(chan uint, 10)

	return &Events{
		header:    title,
		numEvents: 0,
		channel:   ch,
	}, ch
}

// Spin off this function on a new goroutine to
func (ev *Events) Start() {

	fmt.Printf("\n\n")

	for {
		select {
		case val := <-ev.channel:
			{
				if val == UINT_MAX_TERMINATOR {
					break
				} else {
					ev.numEvents += uint64(val)
					ev.update()
				}
			}
		default:
			{
				time.Sleep(time.Microsecond)
			}
		}
	}
}

// update the terminal
func (ev *Events) update() {
	fmt.Printf(
		"%s%s%s\ncount: %v\n",
		PREVIOUS_LINE,
		PREVIOUS_LINE,
		ev.header,
		ev.numEvents,
	)
}

// Send a terminate signal to an event
func StopEvent(ch chan uint) {
	ch <- UINT_MAX_TERMINATOR
	fmt.Println()
}
