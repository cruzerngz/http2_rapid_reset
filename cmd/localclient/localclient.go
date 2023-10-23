package main

import (
	"fmt"
	"http2-rapid-reset/pkg/spec"
)

func main() {
	fmt.Println("Local client")
	var url string = fmt.Sprintf("https://%s:%s/", spec.ServerAddress, spec.ServerPort)
	_ = url
}
