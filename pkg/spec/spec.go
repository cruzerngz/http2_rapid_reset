// Bunch of common stuff placed here
package spec

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// Listening address
var ServerAddress string = ""

// Listening port
var ServerPort string = "4062"

// CtrlcHandler
func CtrlcHandler() {
	channel := make(chan os.Signal)
	signal.Notify(
		channel,
		os.Interrupt,
		syscall.SIGTERM,
	)

	go func() {
		<-channel
		fmt.Printf("\rCTRL-C pressed. Exiting.\n")
		os.Exit(1)
	}()
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "    ")
	return string(s)
}
