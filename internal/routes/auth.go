package routes

import (
	"Astral-Back-End/internal/controllers"

	"github.com/gin-gonic/gin"
)

func AythRouter(router *gin.RouterGroup) {
	groupe := router.Group("/auth")
	{
		groupe.POST("/registration", controllers.Registration)
		groupe.POST("/login", controllers.Login)
		groupe.POST("/refresh", controllers.Refresh)
		groupe.POST("/logout", controllers.Logout)
	}
}
