package httpx

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err != nil {
		errRet := Error(http.StatusBadRequest, err.Error())

		if gStatus, ok := status.FromError(err); ok { // grpc err错误

			errRet = Error(http.StatusBadRequest, gStatus.Message())
		}
		httpx.WriteJson(w, http.StatusBadRequest, errRet)
		return
	}

	ret := Success(resp)
	httpx.WriteJson(w, http.StatusOK, ret)
}
