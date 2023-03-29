package controllers

import (
	"net/http"
	"stage/common"
	"stage/http_server/manager"
	"stage/http_server/request"

	"github.com/gin-gonic/gin"
)

func RedisSet(c *gin.Context) {
	response := common.HttpBack{C: c}
	params := request.KeyValueRequest{}
	err := RequestPostBind(c, &params)
	if err != nil {
		response.Json(http.StatusOK, 1001, nil)
		return
	}
	err = manager.RedisSet(params)
	if err != nil {
		response.Json(http.StatusOK, 1002, nil)
		return
	}
	response.Json(http.StatusOK, 0, nil)
	return
}

func RedisUnset() {

}
