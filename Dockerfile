FROM golang:1.8

RUN go get -v github.com/yireyun/go-queue

EXPOSE 3000

