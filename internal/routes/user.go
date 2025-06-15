package routes

import (
	"Astral-Back-End/internal/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	groupe := router.Group("/user")
	{
		groupe.POST("/", controllers.CreateUser)
	}
}
