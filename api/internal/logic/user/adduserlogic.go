package user

import (
	"context"
	"errors"
	"palworld/rpc/user/pb/user"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var mapRole = map[int]string{
	1: "普通用户",
	2: "管理员",
}

func (l *AddUserLogic) AddUser(req *types.AddUserReq) (*types.AddUserResp, error) {
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("用户名或密码不能为空")
	}
	if mapRole[req.Role] == "" {
		return nil, errors.New("角色不存在")
	}
	in := &user.AddUserReq{
		Username: req.Username,
		Password: req.Password,
		Role:     int32(req.Role),
	}

	resp, err := l.svcCtx.UserRpc.AddUser(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var ret = &types.AddUserResp{
		Id: resp.Id,
	}

	return ret, nil
}
