package main

import (
	"context"
	"fmt"
	"grpc-greeting/greet/greetpb"
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

	client := greetpb.NewGreetServiceClient(conn)

	//doUnary(client)
	//doServerStreaming(client)
	//doClientStreaming(client)
	doBidiStreaming(client)

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

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do an streaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Matías",
			LastName:  "Guerrero",
		},
	}
	resStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling greet many times RPC: %v", err)
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

func doClientStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do an client streaming RPC...")

	requests := []*greetpb.LongGreetRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Matias",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Mario",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Yuly",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Gloria",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Hugo",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Gustavo",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}
	// we iterate over slice requests and send each message individually
	for _, req := range requests {
		fmt.Printf("Sending request: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response form LongGreet: %v", err)
	}
	fmt.Printf("LongGreet response: %v\n", res)
}

func doBidiStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do an Bidi streaming RPC...")

	// Create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating a stream: %v", err)
		return
	}

	requests := []*greetpb.GreetEveryoneRequest{
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Matias",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Mario",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Yuly",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Gloria",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Hugo",
			},
		},
		{
			Greeting: &greetpb.Greeting{
				FirstName: "Gustavo",
			},
		},
	}

	waitc := make(chan struct{})
	// Send a bunch of messages to the client (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	// Receive a bung of messages fron the client (go routine)
	go func() {
		// function to send a bunch of messages
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving from stream: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	// Block until everything is done
	<-waitc
}
