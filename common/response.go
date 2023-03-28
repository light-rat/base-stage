package common

import "github.com/gin-gonic/gin"

type HttpBack struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (h *HttpBack) Json(httpCode int, code int,data interface{}) {
	h.C.JSON(httpCode, Response{
		Code: code,
		Msg:  "",
		Data: data,
	})
	return
}
