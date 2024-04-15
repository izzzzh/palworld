package comments

import (
	"net/http"
	result "palworld/common/httpx"

	"github.com/zeromicro/go-zero/rest/httpx"
	"palworld/api/internal/logic/comments"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
)

func ListCommentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListCommentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := comments.NewListCommentLogic(r.Context(), svcCtx)
		resp, err := l.ListComment(&req)
		result.HttpResult(r, w, resp, err)
	}
}
