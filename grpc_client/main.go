// Package main implements a client for Handler service.
package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "go-grpc-example/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHandlerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.HandleMessage(ctx, &pb.MessageRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not process message: %v", err)
	}
	log.Printf("Reply: %s", r.GetMessage())
}
