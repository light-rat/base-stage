package routes

import (
	"stage/http_server/controllers"

	"github.com/gin-gonic/gin"
)

func Router(c *gin.Engine) {
	testGroup := c.Group("/test")
	{
		testGroup.POST("/redis-set", controllers.RedisSet)
	}

	userGroup := c.Group("/user")
	{
		userGroup.POST("/get-list", controllers.GetUserList)
	}

	homeGroup := c.Group("/home")
	{
		homeGroup.POST("")
	}
}
