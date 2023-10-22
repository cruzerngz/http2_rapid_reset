package main

import (
	"fmt"
	"http2-rapid-reset/pkg/spec"
	"net/http"
)

func main() {

	// transport := &http2.Transport{TLSClientConfig: tlsConfig}
	// _ = transport

	fmt.Println("Client")

	// var client http.Client = *&http.Client{
	// 	Transport: nil,
	// 	CheckRedirect: func(req *http.Request, via []*http.Request) error {
	// 	},
	// 	Jar:     nil,
	// 	Timeout: 0,
	// }

	resp, err := http.Get(fmt.Sprintf("http://%s:%s/example", spec.ServerAddress, spec.ServerPort))
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	fmt.Println(resp)

}
