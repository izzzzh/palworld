package user

import (
	"context"
	"errors"
	"palworld/rpc/user/pb/user"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserReq) (*types.UserInfo, error) {
	if req.Id <= 0 {
		return nil, errors.New("用户ID不能为空")
	}
	in := &user.GetUserInfoReq{
		Id: req.Id,
	}
	resp, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, in)
	if err != nil {
		return nil, err
	}
	var ret = &types.UserInfo{
		Id:        resp.Data.Id,
		Username:  resp.Data.Username,
		Avatar:    resp.Data.Avatar,
		Role:      resp.Data.Role,
		CreatedAt: resp.Data.CreatedAt,
	}

	return ret, nil
}
