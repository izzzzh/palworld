package comments

import (
	"context"
	"palworld/rpc/comments/pb/comments"

	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentLogic {
	return &ListCommentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListCommentLogic) ListComment(req *types.ListCommentReq) (*types.ListCommentResp, error) {

	in := &comments.ListCommentReq{
		Category: req.Category,
		ObjectId: req.ObjectID,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	resp, err := l.svcCtx.CommentsRpc.ListComment(l.ctx, in)
	if err != nil {
		return nil, err
	}
	var ret = &types.ListCommentResp{}

	list := make([]types.Comment, 0)

	for _, v := range resp.Data {
		children := make([]types.Comment, 0)
		for _, child := range v.Children {
			children = append(children, types.Comment{
				ID:         child.Id,
				ParentID:   child.ParentId,
				RootID:     child.RootId,
				UserID:     child.UserId,
				Username:   child.Username,
				Avatar:     child.Avatar,
				Content:    child.Content,
				CreateTime: child.CreatedAt,
			})
		}

		list = append(list, types.Comment{
			ID:         v.Id,
			ParentID:   v.ParentId,
			RootID:     v.RootId,
			UserID:     v.UserId,
			Username:   v.Username,
			Avatar:     v.Avatar,
			Content:    v.Content,
			CreateTime: v.CreatedAt,
			Children:   children,
		})
	}

	ret.List = list
	ret.Total = resp.Total

	return ret, nil
}
