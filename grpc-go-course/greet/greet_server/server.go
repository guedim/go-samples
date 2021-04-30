package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"grpc-greeting/greet/greetpb"
)

type server struct{}

func main() {
	fmt.Println("Hello World")

	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
