package technology_tree

import (
	"net/http"
	result "palworld/common/httpx"

	"palworld/api/internal/logic/technology_tree"
	"palworld/api/internal/svc"
)

func GetTechnologyTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := technology_tree.NewGetTechnologyTreeLogic(r.Context(), svcCtx)
		resp, err := l.GetTechnologyTree()
		result.HttpResult(r, w, resp, err)
	}
}
