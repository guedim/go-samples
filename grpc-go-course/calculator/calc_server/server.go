package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

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

func (*server) PrimeNumberDescomposition(req *calcpb.PrimeRequest, stream calcpb.SumService_PrimeNumberDescompositionServer) error {
	fmt.Printf("PrimeNumberDescomposition function was invoked with %v\n", req)
	k := int32(2)
	N := req.GetPrimeNumber()

	for N > 1 {
		if N%k == 0 {
			res := calcpb.PrimeResponse{
				Result: k,
			}
			stream.Send(&res)
			time.Sleep(500 * time.Millisecond)
			N = N / k
		} else {
			k = k + 1
		}
	}
	return nil
}

func (*server) ComputeAverage(stream calcpb.SumService_ComputeAverageServer) error {
	fmt.Printf("ComputeAverage function was invoked with a streaming request\n")
	sum := 0
	count := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// we have finished reading the client stream
			return stream.SendAndClose(&calcpb.ComputedResponse{
				Result: float32(sum / count),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		number := req.GetNumber()
		sum += int(number)
		count++
	}
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
