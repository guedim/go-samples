package main

import (
	"context"
	"fmt"
	"grpc-greeting/calculator/calcpb"
	"io"
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

	//doUnary(client)

	doServerStreaming(client)

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

func doServerStreaming(c calcpb.SumServiceClient) {
	fmt.Println("Starting to do an streaming RPC...")
	req := &calcpb.PrimeRequest{
		PrimeNumber: 120,
	}
	resStream, err := c.PrimeNumberDescomposition(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling Prime Number Descomposition RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// we have reached end of file
			break
		}
		if err != nil {
			log.Fatalf("error while reading streaming greet many times: %v", err)
		}
		log.Printf("Response from GreetManyTimes: %v", msg.GetResult())
	}
}
