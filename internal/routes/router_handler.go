package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AppRouter(router *gin.Engine) {
	router.GET("/s/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/docs/swagger.json")))
	router.StaticFile("/docs/swagger.json", "./docs/swagger.json")

	groupe := router.Group("/api/v1")
	AythRouter(groupe)
	UserRouter(groupe)
	LingroRouter(groupe)
}
