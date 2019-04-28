// Code generated by protoc-gen-liverpc v0.1, DO NOT EDIT.
// source: v1/RoomEx.proto

package v1

import context "context"

import proto "github.com/golang/protobuf/proto"
import "go-common/library/net/rpc/liverpc"

var _ proto.Message // generate to suppress unused imports

// ================
// RoomEx Interface
// ================

type RoomExRPCClient interface {
	// * 轮播接口
	//
	GetRoundPlayVideo(ctx context.Context, req *RoomExGetRoundPlayVideoReq, opts ...liverpc.CallOption) (resp *RoomExGetRoundPlayVideoResp, err error)
}

// ======================
// RoomEx Live Rpc Client
// ======================

type roomExRPCClient struct {
	client *liverpc.Client
}

// NewRoomExRPCClient creates a client that implements the RoomExRPCClient interface.
func NewRoomExRPCClient(client *liverpc.Client) RoomExRPCClient {
	return &roomExRPCClient{
		client: client,
	}
}

func (c *roomExRPCClient) GetRoundPlayVideo(ctx context.Context, in *RoomExGetRoundPlayVideoReq, opts ...liverpc.CallOption) (*RoomExGetRoundPlayVideoResp, error) {
	out := new(RoomExGetRoundPlayVideoResp)
	err := doRPCRequest(ctx, c.client, 1, "RoomEx.getRoundPlayVideo", in, out, opts)
	if err != nil {
		return nil, err
	}
	return out, nil
}
