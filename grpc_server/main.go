// Package main implements a server for Handler service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "go-grpc-example/message"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// server is used to implement helloworld.HandlerServer.
type server struct {
	pb.UnimplementedHandlerServer
}

// HandleMessage implements message.HandlerServer
func (s *server) HandleMessage(ctx context.Context, in *pb.MessageRequest) (*pb.MessageReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.MessageReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHandlerServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
