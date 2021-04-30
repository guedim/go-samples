package main

import (
	"fmt"
	"grpc-greeting/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello I am a client")

	conn, err := grpc.Dial("localhot:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("created client: %f", c)

}
