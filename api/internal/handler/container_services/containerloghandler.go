package container_services

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"palworld/api/internal/logic/container_services"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"
)

func ContainerLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ContainerLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := container_services.NewContainerLogLogic(r.Context(), svcCtx)
		err := l.ContainerLog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
		w.Header().Set("Cache-Control", "no-transform")
		w.Header().Add("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("X-Accel-Buffering", "no")
		w.Header().Set("Content-Type", "text/event-stream")

		f, ok := w.(http.Flusher)
		if !ok {
			httpx.ErrorCtx(r.Context(), w, errors.New(""))
		}
		w.Header().Set("Cache-Control", "no-transform")
		w.Header().Add("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("X-Accel-Buffering", "no")
		w.Header().Set("Content-Type", "text/event-stream")

		for {
			t, err := json.Marshal("hello world")
			if err != nil {
				break
			}
			if _, err := fmt.Fprintf(w, "data: %s\n\n", t); err != nil {
				break
			}
			f.Flush()
		}
	}
}
