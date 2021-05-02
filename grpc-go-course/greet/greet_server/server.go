package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"grpc-greeting/greet/greetpb"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	fistName := req.GetGreeting().GetFirstName()
	result := "Hello " + fistName

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fmt.Printf("GreetManyTimes function was invoked with %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 20; i++ {
		result := "Hello " + firstName + " number " + strconv.Itoa(i)
		res := greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(&res)
		time.Sleep(500 * time.Millisecond)
	}
	return nil
}

func main() {
	fmt.Println("Starting Greet server ...")

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
