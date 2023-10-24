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

// JSON-stringify a struct.
// Struct needs json tags
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "    ")
	return string(s)
}

// Basic error handler
func HandleError(e error, msg string, exit bool) {
	if e != nil {
		fmt.Printf("error: %s\n", msg)
		fmt.Println(e)

		if exit {
			os.Exit(1)
		}
	}
}
