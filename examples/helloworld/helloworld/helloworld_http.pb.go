// Code generated by protoc-gen-go-http. DO NOT EDIT.

package helloworld

import (
	context "context"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(http.Request)
var _ = new(context.Context)
var _ = binding.MapProto
var _ = mux.NewRouter

const _ = http1.SupportPackageIsVersion1

type GreeterHandler interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

func NewGreeterHandler(srv GreeterHandler, opts ...http1.HandleOption) http.Handler {
	r := mux.NewRouter()

	r.Handle("/helloworld/{name}", http1.NewHandler(srv.SayHello, opts...)).Methods("GET")

	return r
}

type GreeterHttpClient interface {
	SayHello(ctx context.Context, req *HelloRequest, opts ...http1.CallOption) (rsp *HelloReply, err error)
}

type GreeterHttpClientImpl struct {
	cc *http1.Client
}

func NewGreeterHttpClient(client *http1.Client) GreeterHttpClient {
	return &GreeterHttpClientImpl{client}
}

func (c *GreeterHttpClientImpl) SayHello(ctx context.Context, in *HelloRequest, opts ...http1.CallOption) (out *HelloReply, err error) {
	path := "/helloworld/{name}"
	if in != nil {
		path = binding.ProtoPath(path, in)
	}

	out = &HelloReply{}
	err = c.cc.Invoke(ctx, path, in, out, http1.Method("GET"), http1.PathPattern("/helloworld/{name}"), http1.BodyPattern(""), http1.RespBodyPattern(""))
	if err != nil {
		return
	}
	return
}
