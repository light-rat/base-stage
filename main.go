package main

import (
	"stage/http_server/routes"
	"stage/internal/mysql"
	"stage/internal/redis"

	"github.com/gin-gonic/gin"
)

func init() {
	mysql.Init()
	redis.Init()
}

func main() {
	r := gin.Default()
	routes.Router(r)
	r.Run()
}
