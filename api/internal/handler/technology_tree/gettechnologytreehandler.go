package technology_tree

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"palworld/api/internal/logic/technology_tree"
	"palworld/api/internal/svc"
)

func GetTechnologyTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := technology_tree.NewGetTechnologyTreeLogic(r.Context(), svcCtx)
		resp, err := l.GetTechnologyTree()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
