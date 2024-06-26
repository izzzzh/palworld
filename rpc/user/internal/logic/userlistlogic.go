package logic

import (
	"context"
	"palworld/common"
	"palworld/rpc/user/dao"
	"time"

	"palworld/rpc/user/internal/svc"
	"palworld/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var mapRole = map[int]string{
	1: "普通用户",
	2: "管理员",
}

func (l *UserListLogic) UserList(in *user.UserListReq) (*user.UserListResp, error) {
	p := dao.User
	users, err := p.WithContext(l.ctx).Find()
	if err != nil {
		return nil, err
	}
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return nil, err
	}

	var ret = &user.UserListResp{}
	data := make([]*user.User, 0)
	for _, u := range users {
		data = append(data, &user.User{
			Id:        u.ID,
			Username:  u.Name,
			Role:      mapRole[int(u.Role)],
			Avatar:    u.Avatar,
			CreatedAt: u.CreatedAt.In(location).Format(common.TimeLayout),
		})
	}
	ret.Data = data

	return ret, nil
}
