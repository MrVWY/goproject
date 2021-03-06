// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: od.proto

package go_micro_srv_o

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

// Client API for O service

type OService interface {
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type oService struct {
	c    client.Client
	name string
}

func NewOService(name string, c client.Client) OService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.o"
	}
	return &oService{
		c:    c,
		name: name,
	}
}

func (c *oService) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "O.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oService) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "O.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for O service

type OHandler interface {
	Create(context.Context, *Request, *Response) error
	GetAll(context.Context, *Request, *Response) error
}

func RegisterOHandler(s server.Server, hdlr OHandler, opts ...server.HandlerOption) error {
	type o interface {
		Create(ctx context.Context, in *Request, out *Response) error
		GetAll(ctx context.Context, in *Request, out *Response) error
	}
	type O struct {
		o
	}
	h := &oHandler{hdlr}
	return s.Handle(s.NewHandler(&O{h}, opts...))
}

type oHandler struct {
	OHandler
}

func (h *oHandler) Create(ctx context.Context, in *Request, out *Response) error {
	return h.OHandler.Create(ctx, in, out)
}

func (h *oHandler) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.OHandler.GetAll(ctx, in, out)
}
