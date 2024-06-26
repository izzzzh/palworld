package container_services

import (
	"net/http"

	"palworld/api/internal/logic/container_services"
	"palworld/api/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListContainerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := container_services.NewListContainerLogic(r.Context(), svcCtx)
		resp, err := l.ListContainer()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
