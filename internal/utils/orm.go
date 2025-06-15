package utils

import (
	"Astral-Back-End/internal/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ORM(c *gin.Context) *repositories.Repositories {
	repos, exists := c.MustGet("repos").(*repositories.Repositories)
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Repositories not found in context"})
		return nil
	}

	return repos
}
