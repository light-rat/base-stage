package routes

import (
	"github.com/gin-gonic/gin"
	"stage/http_server/controllers"
)

func Router(c *gin.Engine) {
	userGroup := c.Group("/user")
	{
		userGroup.POST("/get-list", controllers.GetUserList)
	}

	homeGroup := c.Group("/home")
	{
		homeGroup.POST("")
	}
}
