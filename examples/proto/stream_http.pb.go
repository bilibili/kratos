// Code generated by protoc-gen-go-http. DO NOT EDIT.

package testproto

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

type StreamServiceHandler interface {
}

func NewStreamServiceHandler(srv StreamServiceHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	return r
}

type StreamServiceHTTPClient interface {
}

type StreamServiceHTTPClientImpl struct {
	cc *http1.Client
}

func NewStreamServiceHTTPClient(client *http1.Client) StreamServiceHTTPClient {
	return &StreamServiceHTTPClientImpl{client}
}
