package post

import (
	proto "github.com/read-in-painless/post-service/proto"
	"context"
)



type Post struct{}

func (g *Post) Create(ctx context.Context, req *proto.PostRequest, rsp *proto.PostReponse) error {
	// rsp.Greeting = "Hello " + req.Name
	return nil
}