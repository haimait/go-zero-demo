// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package usercenter

import (
	"context"

	"go-zero-demo/user-rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetUserInfoReq  = pb.GetUserInfoReq
	GetUserInfoResp = pb.GetUserInfoResp
	User            = pb.User

	Usercenter interface {
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	}

	defaultUsercenter struct {
		cli zrpc.Client
	}
)

func NewUsercenter(cli zrpc.Client) Usercenter {
	return &defaultUsercenter{
		cli: cli,
	}
}

func (m *defaultUsercenter) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := pb.NewUsercenterClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}
