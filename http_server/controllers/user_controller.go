package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stage/common"
	"stage/http_server/manager"
	"stage/http_server/request"
)

func GetUserList(c *gin.Context) {
	response := common.HttpBack{C: c}
	params := request.GetUserListRequest{}
	err := RequestPostBind(c,&params)
	if err != nil {
		response.Json(http.StatusOK, 1001, nil)
		return
	}

	result, err := manager.GetUserList(params)
	response.Json(http.StatusOK, 0, result)
	return
}

func AddUser(c *gin.Context) {
	response := common.HttpBack{C: c}
	_ = response
}
