package main

import (
	"context"
	"fmt"
	"grpc-greeting/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello I am a client")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := greetpb.NewGreetServiceClient(conn)

	doUnary(client)

}

func doUnary(c greetpb.GreetServiceClient) {

	fmt.Println("Starting to do an unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Matias",
			LastName:  "Guerrero",
		},
	}
	resp, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Gret RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", resp.Result)
}
