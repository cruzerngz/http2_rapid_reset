# http2_rapid_reset
Rapid reset demonstrator

## Running locally
Use go versions < `1.20.10` or `1.21.3`, as the vulnerabilities are [patched out](https://groups.google.com/g/golang-announce/c/iNNxDTCjZvo?pli=1) in these versions.

```bash
# start the server in the background

```

## Running in container
```bash
make image
docker compose up
# enter client
docker exec -it client sh
# perform stuff
```
