package main

import (
	"crypto/tls"
	"fmt"
	"http2-rapid-reset/pkg/spec"
	"net/http"
	"time"
)

func main() {

	spec.CtrlcHandler()

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

	transport := http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	var client http.Client = http.Client{
		Transport: &transport,
	}
	_ = client

	var url string = fmt.Sprintf("https://server:%s/", spec.ServerPort)

	// resp, err := http.Post(
	// 	url,
	// 	"json",

	// )
	// if err != nil {
	// 	fmt.Printf("error: %v\n", err)
	// }

	// fmt.Println(resp)

	// Make a GET request.
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	_ = req

	// Send the request.
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("response contents: ", resp.Body)

	for {
		fmt.Println("Waiting")
		time.Sleep(1 * time.Second)
	}
}
