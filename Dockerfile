FROM golang:1.8

# RUN go get -v github.com/yireyun/go-queue
# RUN go get -v github.com/robfig/cron

# RUN go get -v github.com/graphql-go/graphql
# RUN go get -v github.com/golang/protobuf/proto
# RUN go get -v github.com/golang/protobuf/ptypes
 RUN go get -u github.com/gpmgo/gopm
 RUN gopm get -g -d golang.org/x/sys
# RUN gopm get -g -d golang.org/x/text
# RUN gopm get -g -d github.com/golang/net/context
# RUN gopm get -g -d github.com/google/go-genproto
# RUN gopm get -g -d golang.org/x/net/http2

 RUN go get -v github.com/oschwald/geoip2-golang

# grpc下载和部署
# RUN gopm get -g -d github.com/grpc/grpc-go

# RUN mkdir /go/src/google.golang.org
# RUN mkdir /go/src/google.golang.org/grpc
# RUN mkdir /go/src/google.golang.org/genproto
# RUN cp -r /go/src/github.com/grpc/grpc-go/* /go/src/google.golang.org/grpc
# RUN cp -r /go/src/github.com/google/go-genproto/* /go/src/google.golang.org/genproto

# RUN go get -v github.com/davecgh/go-spew/spew
# RUN go get -v github.com/gorilla/mux
# RUN go get -v github.com/joho/godotenv
# RUN go get -v github.com/derekparker/delve/cmd/dlv

EXPOSE 3000
EXPOSE 3001
EXPOSE 50051

