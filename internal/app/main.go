package app

import (
	"Astral-Back-End/internal/repositories"
	"Astral-Back-End/internal/routes"
	"Astral-Back-End/pkg/config"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// @title Astral-Bacl-End API
// @version 1.0
// @description This is a sample server for a Go backend.
// @host localhost:8000
// @BasePath /api/v1

func RunServer() {

	// Initialize server
	server := gin.Default()

	// Initialize config
	config := config.InitConfig()

	// only in production mode
	gin.SetMode(gin.ReleaseMode)

	// Initialize Database
	databse := repositories.RunDB()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://lingro.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, 
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	
	// Initialize middlewares
	server.Use(LoggerMiddleware())                                         // Logger
	server.Use(func(c *gin.Context) { c.Set("repos", databse); c.Next() }) // ORM


	// Initialize routes
	routes.AppRouter(server)

	// Start server
	server.Run(config.Server.Port)
}
