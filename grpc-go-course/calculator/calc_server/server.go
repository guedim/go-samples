package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"grpc-greeting/calculator/calcpb"
)

type server struct {
}

func (*server) Add(ctx context.Context, req *calcpb.SumRequest) (*calcpb.SumResponse, error) {
	fmt.Printf("Add function was invoked with %v\n", req)
	fistNumber := req.GetSum().GetFirstNumber()
	secondNumber := req.GetSum().GetSecondNumber()
	result := fistNumber + secondNumber

	res := &calcpb.SumResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Starting sum server...")

	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calcpb.RegisterSumServiceServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}

}
