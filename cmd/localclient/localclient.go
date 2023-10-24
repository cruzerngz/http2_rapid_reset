// Client used when running locally.
// Make sure that this program is compiled with
// go version < 1.20.10 or < 1.21.3
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
