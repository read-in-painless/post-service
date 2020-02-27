package main

import (
	"fmt"

	micro "github.com/micro/go-micro/v2"
	"github.com/read-in-painless/post-service/post"
	proto "github.com/read-in-painless/post-service/proto"
)

func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(
		micro.Name("post"),
	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterPostHandler(service.Server(), new(post.Post))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
