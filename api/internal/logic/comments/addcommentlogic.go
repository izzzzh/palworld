package comments

import (
	"context"
	"errors"
	"palworld/rpc/comments/pb/comments"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCommentLogic {
	return &AddCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddCommentLogic) AddComment(req *types.AddCommentReq) error {
	if req.Category == "" {
		return errors.New("category is required")
	}
	in := &comments.AddCommentReq{
		Category: req.Category,
		Content:  req.Content,
		RootId:   req.RootID,
		ParentId: req.ParentID,
		UserId:   req.UserID,
		ObjectId: req.ObjectID,
	}

	_, err := l.svcCtx.CommentsRpc.AddComment(l.ctx, in)
	if err != nil {
		return err
	}

	return nil
}
