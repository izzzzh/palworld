package logic

import (
	"context"
	"errors"
	"net/http"
	"palworld/common"
	"palworld/rpc/comments/dao"
	"palworld/rpc/comments/internal/svc"
	"palworld/rpc/comments/pb/comments"
	"palworld/rpc/user/pb/user"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentLogic {
	return &ListCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListCommentLogic) ListComment(in *comments.ListCommentReq) (*comments.CommentListResp, error) {
	p := dao.Comments
	page, pageSize := int(in.Page), int(in.PageSize)
	if in.PageSize <= 0 {
		pageSize = 10
	}
	if in.Page <= 0 {
		page = 1
	}
	offset := (page - 1) * pageSize
	commentsResp, err := p.Where(p.Category.Eq(in.Category)).
		Where(p.ObjectID.Eq(in.ObjectId)).
		Limit(pageSize).Offset(offset).
		Where(p.RootCommentID.Eq(0)).
		Order(p.CreatedAt.Desc()).Find()
	if err != nil {
		return nil, err
	}
	var commentIds []int64
	for _, comment := range commentsResp {
		commentIds = append(commentIds, comment.ID)
	}

	var userIds []int64
	err = p.Where(p.RootCommentID.In(commentIds...)).
		Or(p.ID.In(commentIds...)).Distinct(p.UserID).Select(p.UserID).Scan(&userIds)
	if err != nil {
		return nil, err
	}

	usersRet, err := l.svcCtx.UserRpc.GetUsersByIds(l.ctx, &user.GetUsersByIdsReq{Ids: userIds})
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]*user.User)
	for _, val := range usersRet.Data {
		userMap[val.Id] = val
	}

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return nil, err
	}

	data := make([]*comments.Comment, 0)
	for _, comment := range commentsResp {
		children, err := p.Where(p.RootCommentID.Eq(comment.ID)).Find()
		if err != nil {
			return nil, err
		}
		childComments := make([]*comments.Comment, 0)
		for _, child := range children {
			childComments = append(childComments, &comments.Comment{
				Id:        child.ID,
				UserId:    child.UserID,
				Username:  userMap[child.UserID].Username,
				Avatar:    userMap[child.UserID].Avatar,
				Content:   child.Content,
				CreatedAt: comment.CreatedAt.In(location).Format(common.TimeLayout),
				RootId:    child.RootCommentID,
				ParentId:  child.ParentCommentID,
			})
		}
		data = append(data, &comments.Comment{
			Id:        comment.ID,
			UserId:    comment.UserID,
			Username:  userMap[comment.UserID].Username,
			Avatar:    userMap[comment.UserID].Avatar,
			Content:   comment.Content,
			CreatedAt: comment.CreatedAt.In(location).Format(common.TimeLayout),
			Children:  childComments,
		})
	}

	var ret = &comments.CommentListResp{}
	ret.Code = http.StatusOK
	ret.Message = "ok"
	ret.Data = data

	return ret, nil
}

func (l *ListCommentLogic) validate(in *comments.ListCommentReq) error {
	if in.Category == "" {
		return errors.New("category is empty")
	}
	if _, ok := AllowCategory[in.Category]; !ok {
		return errors.New("category not allow")
	}
	if in.ObjectId <= 0 {
		return errors.New("object_id is empty")
	}

	return nil
}
