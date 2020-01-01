// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/post.proto

package post

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Post service

type PostService interface {
	Create(ctx context.Context, in *PostRequest, opts ...client.CallOption) (*PostReponse, error)
}

type postService struct {
	c    client.Client
	name string
}

func NewPostService(name string, c client.Client) PostService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "post"
	}
	return &postService{
		c:    c,
		name: name,
	}
}

func (c *postService) Create(ctx context.Context, in *PostRequest, opts ...client.CallOption) (*PostReponse, error) {
	req := c.c.NewRequest(c.name, "Post.Create", in)
	out := new(PostReponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Post service

type PostHandler interface {
	Create(context.Context, *PostRequest, *PostReponse) error
}

func RegisterPostHandler(s server.Server, hdlr PostHandler, opts ...server.HandlerOption) error {
	type post interface {
		Create(ctx context.Context, in *PostRequest, out *PostReponse) error
	}
	type Post struct {
		post
	}
	h := &postHandler{hdlr}
	return s.Handle(s.NewHandler(&Post{h}, opts...))
}

type postHandler struct {
	PostHandler
}

func (h *postHandler) Create(ctx context.Context, in *PostRequest, out *PostReponse) error {
	return h.PostHandler.Create(ctx, in, out)
}