FROM golang:1.8

RUN go get -v github.com/yireyun/go-queue
RUN go get -v github.com/robfig/cron
# RUN go get -v github.com/golang/net/context
RUN go get -v google.golang.org/grpc

RUN go get -v github.com/davecgh/go-spew/spew
RUN go get -v github.com/gorilla/mux
RUN go get -v github.com/joho/godotenv
RUN go get -v github.com/derekparker/delve/cmd/dlv

EXPOSE 3000
EXPOSE 3001

