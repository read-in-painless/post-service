package post

import (
	"context"
	proto "github.com/read-in-painless/post-service/proto"
)

type Post struct{}

func (g *Post) Create(ctx context.Context, req *proto.PostRequest, rsp *proto.PostReponse) error {
	rsp.Status = proto.Status_ACCEPTED
	rsp.Message = "sucess"
	return nil
}
