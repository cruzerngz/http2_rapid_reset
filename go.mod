module http2-rapid-reset

// we specifically want to avoid versions >= 1.21.3 or 1.20.10,
// because these versions have patched out the exploit.
// When performing a multistage docker build, remember to use
// a build image that does not have the above go versions!
//
// https://groups.google.com/g/golang-announce/c/iNNxDTCjZvo
go 1.21

require (
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
