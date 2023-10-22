package main

import (
	"errors"
	"fmt"
	"http2-rapid-reset/pkg/spec"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/net/http2"
)

func main() {

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/example", getExample)

	fmt.Println("Hello server")
	serverAddress := fmt.Sprintf("%s:%s", spec.ServerAddress, spec.ServerPort)

	var server http.Server = http.Server{
		Addr: serverAddress,
	}

	var err error

	// make server support http2 requests
	err = http2.ConfigureServer(&server, &http2.Server{})
	handleError(err, "unable to configure server to http2", true)

	fmt.Printf("Server listening at %s\n", serverAddress)

	executable, err := os.Executable()
	handleError(err, "unable to get executable path", true)

	exDir := filepath.Dir(executable)
	err = server.ListenAndServeTLS(fmt.Sprintf("%s/server.crt", exDir), fmt.Sprintf("%s/server.key", exDir))

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	}
	handleError(err, "server error", true)
}

// Basic error handler
func handleError(e error, msg string, exit bool) {
	if e != nil {
		fmt.Printf("error: %s\n", msg)
		fmt.Println(e)

		if exit {
			os.Exit(1)
		}
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Received request")
	fmt.Printf("%#v\n", r)
	io.WriteString(w, "Welcome to the demo server!")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getExample(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /exmple request\n")
	fmt.Printf("%#v\n", r)
	io.WriteString(w, "Hello, HTTP!\n")
}
