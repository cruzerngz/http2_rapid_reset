# docker build -t build -f docker/build.Dockerfile .
FROM golang:1.21.2-bullseye as BUIDER

WORKDIR /build
COPY . .

# create a statically linked binary
ENV CGO_ENABLED=0

RUN make keys
RUN make build
