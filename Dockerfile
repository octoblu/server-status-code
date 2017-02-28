FROM golang:1.8
MAINTAINER Octoblu, Inc. <docker@octoblu.com>

WORKDIR /go/src/github.com/octoblu/server-status-code
COPY . /go/src/github.com/octoblu/server-status-code

RUN env CGO_ENABLED=0 go build -o server-status-code -a -ldflags '-s' .

CMD ["./server-status-code"]
