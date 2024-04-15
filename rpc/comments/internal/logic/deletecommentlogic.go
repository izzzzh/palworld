package logic

import (
	"context"

	"palworld/rpc/comments/internal/svc"
	"palworld/rpc/comments/pb/comments"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteCommentLogic) DeleteComment(in *comments.DeleteCommentReq) (*comments.DeleteCommentResp, error) {
	// todo: add your logic here and delete this line

	return &comments.DeleteCommentResp{}, nil
}
