package controllers

import "github.com/gin-gonic/gin"

func RequestPostBind(ctx *gin.Context, opt interface{}) error{
	err := ctx.BindJSON(opt)
	return err
}

func RequestGetBind(ctx *gin.Context, opt interface{}) error{
	err := ctx.BindQuery(opt)
	return err
}