package main

import (
	"log"
	"net"

//	pb "github.com/grpc/grpc-go/examples/helloworld/helloworld"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
//	"github.com/grpc/grpc-go"
)

const (
	port  =  ":50051"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest)(*pb.HelloReply, error){
	log.Println(in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main(){
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}