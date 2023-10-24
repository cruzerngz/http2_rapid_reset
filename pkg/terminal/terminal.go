// For manipulating terminal stuffs

package terminal

import (
	"fmt"
	"time"
)

const UINT_MAX_TERMINATOR uint = ^uint(0)
const PREVIOUS_LINE string = "\033[F"
const CLEAR_LINE string = "\033[2K\r"

// Events tally counter
type Events struct {
	header    string
	numEvents uint64
	// scratch data for previous occurence
	prevTime   time.Time
	prevEvents uint64
	channel    chan uint
}

// Create a new event tracker along with its attached channel
func NewEvent(title string) (*Events, chan uint) {

	var ch chan uint = make(chan uint, 10)

	return &Events{
		header:     title,
		numEvents:  1,
		channel:    ch,
		prevTime:   time.Now(),
		prevEvents: 0,
	}, ch
}

// Spin off this function on a new goroutine to
func (ev *Events) Start() {

	fmt.Printf("\n\n\n")

	for {
		select {
		case val := <-ev.channel:
			{
				curTime := time.Now()
				delta := curTime.Sub(ev.prevTime)
				ev.prevTime = curTime

				deltaEv := ev.numEvents - ev.prevEvents
				ev.prevEvents = ev.numEvents

				if val == UINT_MAX_TERMINATOR {
					break
				} else {
					ev.numEvents += uint64(val)
					ev.update(delta / time.Duration(deltaEv))
				}
			}
		default:
			{
				// yield as little time as possible
				time.Sleep(time.Microsecond)
			}
		}
	}
}

// update the terminal
func (ev *Events) update(delta time.Duration) {

	var freq float64 = 1.0 / delta.Seconds()

	fmt.Printf(
		"%s%s%s%s\ncount: %v\n%sfrequency: %.02f/s\n",
		PREVIOUS_LINE,
		PREVIOUS_LINE,
		PREVIOUS_LINE,
		ev.header,
		ev.numEvents,
		CLEAR_LINE,
		freq,
	)
}

// Send a terminate signal to an event
func StopEvent(ch chan uint) {
	ch <- UINT_MAX_TERMINATOR
	fmt.Println()
}
