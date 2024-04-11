// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userserver

import (
	"context"

	"palworld/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddUserReq       = user.AddUserReq
	AddUserResp      = user.AddUserResp
	GetCaptchaReq    = user.GetCaptchaReq
	GetCaptchaResp   = user.GetCaptchaResp
	GetUserInfoReq   = user.GetUserInfoReq
	GetUserInfoResp  = user.GetUserInfoResp
	LoginInfo        = user.LoginInfo
	LoginReq         = user.LoginReq
	LoginResp        = user.LoginResp
	RegisterUserReq  = user.RegisterUserReq
	RegisterUserResp = user.RegisterUserResp
	Token            = user.Token
	User             = user.User
	UserListReq      = user.UserListReq
	UserListResp     = user.UserListResp

	UserServer interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error)
		GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaResp, error)
		RegisterUser(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*RegisterUserResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	}

	defaultUserServer struct {
		cli zrpc.Client
	}
)

func NewUserServer(cli zrpc.Client) UserServer {
	return &defaultUserServer{
		cli: cli,
	}
}

func (m *defaultUserServer) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := user.NewUserServerClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserServer) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := user.NewUserServerClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}

func (m *defaultUserServer) AddUser(ctx context.Context, in *AddUserReq, opts ...grpc.CallOption) (*AddUserResp, error) {
	client := user.NewUserServerClient(m.cli.Conn())
	return client.AddUser(ctx, in, opts...)
}

func (m *defaultUserServer) GetCaptcha(ctx context.Context, in *GetCaptchaReq, opts ...grpc.CallOption) (*GetCaptchaResp, error) {
	client := user.NewUserServerClient(m.cli.Conn())
	return client.GetCaptcha(ctx, in, opts...)
}

func (m *defaultUserServer) RegisterUser(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*RegisterUserResp, error) {
	client := user.NewUserServerClient(m.cli.Conn())
	return client.RegisterUser(ctx, in, opts...)
}

func (m *defaultUserServer) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := user.NewUserServerClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}
