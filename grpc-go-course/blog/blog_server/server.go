package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"

	"grpc-greeting/blog/blogpb"
)

var collection *mongo.Collection

type server struct{}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitemty"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
	Title    string             `bson:"title"`
}

func main() {

	// if we creash the go code, we get the file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Connect to MongoDB
	fmt.Println("Connecting to MongoDB ...")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	collection = client.Database("myGoDb").Collection("blog")

	fmt.Println("Starting Blog server ...")
	list, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	s := grpc.NewServer(opts...)
	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {
		fmt.Println("Starting Server...")
		if err := s.Serve(list); err != nil {
			log.Fatalf("Failed to server: %v", err)
		}
	}()

	// Wait for Control C for exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	fmt.Println("Stopping the server")
	s.Stop()
	fmt.Println("Stopping the listener")
	list.Close()
	fmt.Println("Closing MogoDb connection")
	client.Disconnect(context.TODO())
	fmt.Println("End of Program")

}
