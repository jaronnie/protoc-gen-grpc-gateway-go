// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package user

import (
	"context"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"

	"github.com/jaronnie/protoc-gen-grpc-gateway-gosdk/grpc-restful/userpb"
)

type (
	AddUserReq  = userpb.AddUserReq
	AddUserResp = userpb.AddUserResp

	User interface {
		Add(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Add(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error) {
	client := userpb.NewUserClient(m.cli.Conn())
	return client.Add(ctx, in, opts...)
}
