package comments

import (
	"net/http"
	result "palworld/common/httpx"

	"github.com/zeromicro/go-zero/rest/httpx"
	"palworld/api/internal/logic/comments"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
)

func AddCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comments.NewAddCommentLogic(r.Context(), svcCtx)
		err := l.AddComment(&req)
		result.HttpResult(r, w, struct{}{}, err)
	}
}
