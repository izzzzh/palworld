// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"palworld/rpc/user/internal/logic"
	"palworld/rpc/user/internal/svc"
	"palworld/rpc/user/pb/user"
)

type UserServerServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServerServer
}

func NewUserServerServer(svcCtx *svc.ServiceContext) *UserServerServer {
	return &UserServerServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServerServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServerServer) UserList(ctx context.Context, in *user.UserListReq) (*user.UserListResp, error) {
	l := logic.NewUserListLogic(ctx, s.svcCtx)
	return l.UserList(in)
}

func (s *UserServerServer) AddUser(ctx context.Context, in *user.AddUserReq) (*user.AddUserResp, error) {
	l := logic.NewAddUserLogic(ctx, s.svcCtx)
	return l.AddUser(in)
}

func (s *UserServerServer) GetCaptcha(ctx context.Context, in *user.GetCaptchaReq) (*user.GetCaptchaResp, error) {
	l := logic.NewGetCaptchaLogic(ctx, s.svcCtx)
	return l.GetCaptcha(in)
}

func (s *UserServerServer) RegisterUser(ctx context.Context, in *user.RegisterUserReq) (*user.RegisterUserResp, error) {
	l := logic.NewRegisterUserLogic(ctx, s.svcCtx)
	return l.RegisterUser(in)
}

func (s *UserServerServer) GetUserInfo(ctx context.Context, in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}
