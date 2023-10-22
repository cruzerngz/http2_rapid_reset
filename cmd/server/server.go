package main

import (
	"errors"
	"fmt"
	"http2-rapid-reset/pkg/spec"
	"io"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/example", getExample)

	fmt.Println("Hello server")

	server_add := fmt.Sprintf("%s:%s", spec.ServerAddress, spec.ServerPort)
	// fmt.Printf("Server listening at %s\n", server_add)

	err := http.ListenAndServe(server_add, nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Server closed")
	}
	if err != nil {
		fmt.Printf("server start error: %v\n", err)
		os.Exit(1)
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
