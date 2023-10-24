package main

import (
	"fmt"
	"http2-rapid-reset/pkg/receiver"
	"http2-rapid-reset/pkg/spec"
	"os"
	"path/filepath"
)

func main() {

	spec.CtrlcHandler()

	var err error

	// http.HandleFunc("/", getRoot)
	// http.HandleFunc("/hello", getHello)
	// http.HandleFunc("/example", getExample)

	fmt.Println("Hello server")
	serverAddress := fmt.Sprintf("%s:%s", spec.ServerAddress, spec.ServerPort)

	executable, err := os.Executable()
	spec.HandleError(err, "failed to get executable info", true)

	exDir := filepath.Dir(executable)
	var crt string = fmt.Sprintf("%s/server.crt", exDir)
	var key string = fmt.Sprintf("%s/server.key", exDir)

	s, err := receiver.NewH2Server(
		serverAddress,
		crt,
		key,
	)

	spec.HandleError(err, "Failed to initialize server", true)

	err = s.Start()
	spec.HandleError(err, "Error starting server", true)

}

// func getRoot(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("Received request")
// 	fmt.Printf("%#v\n", r)
// 	io.WriteString(w, "Welcome to the demo server!")
// }

// func getHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got /hello request\n")
// 	io.WriteString(w, "Hello, HTTP!\n")
// }

// func getExample(w http.ResponseWriter, r *http.Request) {
// 	fmt.Printf("got /exmple request\n")
// 	fmt.Printf("%#v\n", r)
// 	io.WriteString(w, "Hello, HTTP!\n")
// }
