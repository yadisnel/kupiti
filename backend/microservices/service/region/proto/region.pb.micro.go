// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/region.proto

package region

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Region service

func NewRegionEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Region service

type RegionService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Region_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Region_PingPongService, error)
}

type regionService struct {
	c    client.Client
	name string
}

func NewRegionService(name string, c client.Client) RegionService {
	return &regionService{
		c:    c,
		name: name,
	}
}

func (c *regionService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Region.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Region_StreamService, error) {
	req := c.c.NewRequest(c.name, "Region.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &regionServiceStream{stream}, nil
}

type Region_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type regionServiceStream struct {
	stream client.Stream
}

func (x *regionServiceStream) Close() error {
	return x.stream.Close()
}

func (x *regionServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *regionServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *regionServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *regionServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *regionService) PingPong(ctx context.Context, opts ...client.CallOption) (Region_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Region.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &regionServicePingPong{stream}, nil
}

type Region_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type regionServicePingPong struct {
	stream client.Stream
}

func (x *regionServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *regionServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *regionServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *regionServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *regionServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *regionServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Region service

type RegionHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Region_StreamStream) error
	PingPong(context.Context, Region_PingPongStream) error
}

func RegisterRegionHandler(s server.Server, hdlr RegionHandler, opts ...server.HandlerOption) error {
	type region interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Region struct {
		region
	}
	h := &regionHandler{hdlr}
	return s.Handle(s.NewHandler(&Region{h}, opts...))
}

type regionHandler struct {
	RegionHandler
}

func (h *regionHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.RegionHandler.Call(ctx, in, out)
}

func (h *regionHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RegionHandler.Stream(ctx, m, &regionStreamStream{stream})
}

type Region_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type regionStreamStream struct {
	stream server.Stream
}

func (x *regionStreamStream) Close() error {
	return x.stream.Close()
}

func (x *regionStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *regionStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *regionStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *regionStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *regionHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.RegionHandler.PingPong(ctx, &regionPingPongStream{stream})
}

type Region_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type regionPingPongStream struct {
	stream server.Stream
}

func (x *regionPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *regionPingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *regionPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *regionPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *regionPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *regionPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}