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
		log.Fatalf("Unexpected error creating blog:  %v \n", err)
	}
	fmt.Printf("Blog has been created: %v\n", createBlogRes)
	blogId := createBlogRes.GetBlog().GetId()

	// Create blog
	fmt.Println("Reading a blog by id")
	_, err2 := client.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "609a03c4a3d154328913d451"})
	if err2 != nil {
		fmt.Printf("Error finding a blog by id: %v\n", err2)
	}
	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogId}
	readBlogRes, err3 := client.ReadBlog(context.Background(), readBlogReq)
	if err3 != nil {
		fmt.Printf("Error finding a blog by id: %v\n", err3)
	}
	fmt.Printf("Blog read with id: %v \n", readBlogRes)

	// update blog
	fmt.Println("updating  the blog")
	newBlog := &blogpb.Blog{
		Id:       blogId,
		AuthorId: "Matías (edited)",
		Title:    "Matines blog (edited)",
		Content:  "Matines day to day blog (edited)",
	}
	updateRes, err4 := client.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if err4 != nil {
		fmt.Printf("Error updating  a blog by id: %v\n", err4)
	}
	fmt.Printf("Blog was updated : %v \n", updateRes)

}
