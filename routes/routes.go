package routes

import (
	"doc/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	home := router.Group("/")
	{
		home.GET("", controllers.HomeController)
	}

	user := router.Group("/user")
	{
		user.GET("", controllers.UserController)
	}
}