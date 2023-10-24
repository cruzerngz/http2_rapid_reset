// Server implementations
package receiver

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// Struct for attached methods
type Server struct {
	// addr  string
	cert  string
	key   string
	inner *http.Server
}

// Create a http2 server
func NewH2Server(addr string, cert string, key string) (*Server, error) {

	handler := http.HandlerFunc(serverHandler)

	h2s := http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(handler, &http2.Server{}),
	}

	err := http2.ConfigureServer(&h2s, &http2.Server{})
	if err != nil {
		return nil, err
	}

	var s Server = Server{
		cert:  cert,
		key:   key,
		inner: &h2s,
	}

	return &s, nil
}

// Start the server
func (server *Server) Start() error {

	err := server.inner.ListenAndServeTLS(
		server.cert,
		server.key,
	)

	return err
}

// Main handler, invoked on server request
func serverHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Demo http2 server response\n")

	fmt.Println(r.Proto)
	t := time.Now()
	fmt.Printf("[%v] Received %v request on %v with length %v\n",
		t,
		r.Method,
		r.URL,
		r.ContentLength,
	)

	// spew.Dump(&r)

}
