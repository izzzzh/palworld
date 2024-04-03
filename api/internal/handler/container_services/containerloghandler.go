package container_services

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"palworld/api/internal/logic/container_services"
	"palworld/api/internal/svc"
	"palworld/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ContainerLogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ContainerLogReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		f, ok := w.(http.Flusher)
		if !ok {
			httpx.ErrorCtx(r.Context(), w, errors.New(""))
		}

		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("X-Accel-Buffering", "no")
		w.Header().Set("Content-Type", "text/event-stream")

		l := container_services.NewContainerLogLogic(r.Context(), svcCtx)
		reader, err := l.ContainerLog(&req)
		if err != nil {
			if err == io.EOF {
				fmt.Fprintf(w, "event: container-stopped\ndata: end of stream\n\n")
				f.Flush()
			} else {
				httpx.ErrorCtx(r.Context(), w, err)
			}
			return
		}
		defer reader.Close()

		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		go func() {
			for {
				select {
				case <-r.Context().Done():
					return
				case <-ticker.C:
					fmt.Fprintf(w, ":ping \n\n")
					f.Flush()
				}
			}
		}()

		buffered := bufio.NewReader(reader)
		eventGen := container_services.NewEventGenerator(buffered)

		for {
			logEvent := eventGen.Next()
			if logEvent != nil {
				if buf, err := json.Marshal(logEvent); err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Fprintf(w, "data: %s\n", buf)
				}
				if logEvent.Timestamp > 0 {
					fmt.Fprintf(w, "id: %d\n", logEvent.Timestamp)
				}
				fmt.Fprintf(w, "\n")
				f.Flush()
			}
		}
	}
}
