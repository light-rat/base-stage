package main

import (
	"github.com/gin-gonic/gin"
	"stage/http_server/routes"
	"stage/internal/mysql"
)

func init() {
	mysql.Init()
}

func main() {
	r := gin.Default()
	routes.Router(r)
	r.Run()
}