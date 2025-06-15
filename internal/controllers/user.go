package controllers

import (
	"Astral-Back-End/internal/repositories"
	"Astral-Back-End/internal/services"
	"Astral-Back-End/internal/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Create a new user
// @Description Simple creating user
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {User created successfully}
// @Router /user [post]
func CreateUser(context *gin.Context) {
	var newUser repositories.User
	db := utils.ORM(context)
	UserService := services.NewUserService(db.UserRepo)

	if err := context.ShouldBindJSON(&newUser); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := UserService.CreateUser(&newUser); err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "User created successfully"})
}
