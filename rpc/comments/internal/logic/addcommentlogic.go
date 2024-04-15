package logic

import (
	"context"
	"errors"
	"palworld/rpc/comments/dao"
	"palworld/rpc/comments/model"
	"palworld/rpc/user/pb/user"

	"palworld/rpc/comments/internal/svc"
	"palworld/rpc/comments/pb/comments"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var (
	AllowCategory = map[string]struct{}{"pal": {}, "skill": {}, "goods": {}}
)

func (l *AddCommentLogic) AddComment(in *comments.AddCommentReq) (*comments.AddCommentResp, error) {
	if _, ok := AllowCategory[in.Category]; !ok {
		return nil, errors.New("不允许的分类")
	}

	if _, err := l.svcCtx.UserRpc.GetUserInfo(l.ctx, &user.GetUserInfoReq{Id: in.UserId}); err != nil {
		return nil, err
	}

	p := dao.Comments
	comment := &model.Comments{
		UserID:          in.UserId,
		Content:         in.Content,
		RootCommentID:   in.RootId,
		ParentCommentID: in.ParentId,
		Category:        in.Category,
		ObjectID:        in.ObjectId,
	}

	err := p.Create(comment)
	if err != nil {
		return nil, err
	}

	var ret = &comments.AddCommentResp{}
	ret.Id = comment.ID

	return ret, nil
}
