package main

import (
	"context"
	"fmt"
	"grpc-greeting/calculator/calcpb"
	"io"
	"log"
	"time"

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
	//doServerStreaming(client)
	doClientStreaming(client)

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

func doClientStreaming(c calcpb.SumServiceClient) {
	fmt.Println("Starting to do a compute average client streaming RPC...")

	requests := []*calcpb.NumberRequest{
		{
			Number: 1,
		},
		{
			Number: 2,
		},
		{
			Number: 3,
		},
		{
			Number: 4,
		},
	}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling ComputeAverage: %v", err)
	}
	// we iterate over slice requests and send each message individually
	for _, req := range requests {
		fmt.Printf("Sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response form ComputeAverage: %v", err)
	}
	fmt.Printf("ComputeAverage response: %v\n", res.GetAverage())
}
