FROM golang:1.8

RUN go get -v github.com/yireyun/go-queue
RUN go get -v github.com/robfig/cron

EXPOSE 3000

