# http2_rapid_reset
Rapid reset demonstrator

## Running locally
Use go versions < `1.20.10` or `1.21.3`, as the vulnerabilities are [patched out](https://groups.google.com/g/golang-announce/c/iNNxDTCjZvo?pli=1) in these versions.

```bash
make keys
make build

# start the server in the background, or another terminal
./bin/server &
./bin/localclient -duration 10s -frequency 10000
```

## Running in container
```bash
make image
docker compose up
# enter client
docker exec -it client sh
# perform stuff
client -duration 10s -frequency 10000 # rapid reset example
ddosclient -duration 10s -frequency 10000 # normal ddos example
```

## Explanation
The HTTP2 rapid reset attack exploits the behaviour of servers that support HTTP/2.

The client container sends out a predetermined volume of requests to a HTTP2 server,
sent as a stream of 100 (default) requests at a time. This stream is terminated by the client before the request completes.

This imparts little computational work client-side, while the server still needs to
dedicate resources to processing these cancelled requests.
