package main

import (
	"context"
	"fmt"
	"grpc-greeting/blog/blogpb"
	"log"

	"google.golang.org/grpc"
)

func main() {

	fmt.Println("Hello I am a blog client")

	opts := grpc.WithInsecure()
	conn, err := grpc.Dial("localhost:50051", opts)
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	client := blogpb.NewBlogServiceClient(conn)

	// Create blog
	fmt.Println("Creating the blog")
	blog := &blogpb.Blog{
		AuthorId: "Matías",
		Title:    "Matines blog",
		Content:  "Matines day to day blog",
	}

	createBlogRes, err := client.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: blog})
	if err != nil {
		log.Fatalf("Unexpected error creating blog:  %v", err)
	}
	fmt.Printf("Blog has been created: %v", createBlogRes)
}
