package logic

import (
	"context"
	"errors"
	"net/http"
	"palworld/rpc/user/dao"
	"palworld/rpc/user/model"

	"palworld/rpc/user/internal/svc"
	"palworld/rpc/user/pb/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var (
	DefaultAvatar = "https://ppcat.fun/palworld/images/pals/SheepBall.png"
)

func (l *AddUserLogic) AddUser(in *user.AddUserReq) (*user.AddUserResp, error) {
	if in.Username == " " || in.Password == " " {
		return nil, errors.New("用户名或密码不能为空")
	}

	enPassword, err := EncryptPassword(in.Password)
	if err != nil {
		return nil, err
	}
	u := &model.User{
		Name:     in.Username,
		Password: enPassword,
		Role:     in.Role,
		Avatar:   DefaultAvatar,
	}
	if err = AddUser(l.ctx, u); err != nil {
		return nil, err
	}
	var ret = &user.AddUserResp{}
	ret.Code = http.StatusOK
	ret.Message = "添加成功"
	ret.Id = u.ID

	return ret, nil
}

func VerifyUserExist(ctx context.Context, username string) (bool, error) {
	p := dao.User
	loginUser, err := p.WithContext(ctx).Where(p.Name.Eq(username)).First()
	if err != nil {
		return false, err
	}

	return loginUser.ID > 0, nil
}

func AddUser(ctx context.Context, u *model.User) error {
	p := dao.User
	err := p.WithContext(ctx).Create(u)
	if err != nil {
		return err
	}
	return nil
}
