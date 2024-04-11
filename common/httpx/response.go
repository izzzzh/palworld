package httpx

import "net/http"

type ResponseSuccessBean struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type NullJson struct{}

func Success(data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{http.StatusOK, "OK", data}
}

type ResponseErrorBean struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Error(errCode int, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg}
}
