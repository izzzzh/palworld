package user

import (
	"context"
	"palworld/rpc/user/pb/user"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserListLogic {
	return &GetUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserListLogic) GetUserList(req *types.UserListReq) (*types.UserListResp, error) {
	in := &user.UserListReq{}

	resp, err := l.svcCtx.UserRpc.UserList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var ret = &types.UserListResp{}
	userList := make([]*types.UserInfo, 0)
	for _, val := range resp.Data {
		userList = append(userList, &types.UserInfo{
			Id:        val.Id,
			Username:  val.Username,
			Avatar:    val.Avatar,
			Role:      val.Role,
			CreatedAt: val.CreatedAt,
		})
	}

	return ret, nil
}
