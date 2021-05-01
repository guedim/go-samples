package main

import (
	"context"
	"fmt"
	"grpc-greeting/calculator/calcpb"
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

	client := calcpb.NewSumServiceClient(conn)

	doUnary(client)

}

func doUnary(c calcpb.SumServiceClient) {

	fmt.Println("Starting to do an unary RPC...")
	req := &calcpb.SumRequest{
		Sum: &calcpb.Sum{
			FirstNumber:  7,
			SecondNumber: 8,
		},
	}
	resp, err := c.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling Add RPC: %v", err)
	}

	log.Printf("Response from Greet: %v", resp.Result)
}
