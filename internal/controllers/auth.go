package controllers

import (
	"Astral-Back-End/internal/repositories"
	"Astral-Back-End/internal/services"
	"Astral-Back-End/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// @Summary Registration
// @Description Registration with email
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body repositories.User true "User data"
// @Success 200 {object} map[string]string "message: User created successfully"
// @Router /auth/registration [post]
func Registration(context *gin.Context) {
	var newUser repositories.User
	db := utils.ORM(context)
	UserService := services.NewUserService(db.UserRepo)
	AuthService := services.NewAuthService(UserService)

	if err := context.ShouldBindJSON(&newUser); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := AuthService.Registrarion(&newUser); err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "User created successfully"})
}

// @Summary Login
// @Description Login with email
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "User data"
// @Success 200 {object} map[string]string "message: User created successfully"
// @Router /auth/login [post]
func Login(context *gin.Context) {
	var login LoginRequest

	db := utils.ORM(context)
	UserService := services.NewUserService(db.UserRepo)
	AuthService := services.NewAuthService(UserService)

	if err := context.ShouldBindJSON(&login); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, user, err := AuthService.Login(login.Email, login.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	setTokenToCookie(context, refreshToken)

	context.JSON(http.StatusOK, gin.H{
		"tokens": gin.H{
			"access_token": accessToken,
		},
		"user": UserResponse{
			ID:    int(user.Id),
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

// @Summary Refresh Tokens
// @Description Refresh access token using refresh token from cookie
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "access_token: New access token"
// @Failure 401 {object} map[string]string "error: Invalid or expired refresh token"
// @Router /auth/refresh [post]
func Refresh(context *gin.Context) {
	// Получаем refresh-токен из куки
	refreshToken, err := context.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token not found in cookie"})
		return
	}

	// Инициализируем сервисы
	db := utils.ORM(context)
	UserService := services.NewUserService(db.UserRepo)
	AuthService := services.NewAuthService(UserService)

	// Обновляем токены
	newAccessToken, newRefreshToken, user, err := AuthService.Refresh(refreshToken)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Устанавливаем новый refresh-токен в куки
	setTokenToCookie(context, newRefreshToken)

	context.JSON(http.StatusOK, gin.H{
		"tokens": gin.H{
			"access_token": newAccessToken,
		},
		"user": UserResponse{
			ID:    int(user.Id),
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

// @Summary Logout
// @Description Logout user by clearing refresh token cookie
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string "message: Successfully logged out"
// @Router /auth/logout [post]
func Logout(context *gin.Context) {
	// Очищаем refresh-токен в куки, устанавливая его срок действия в прошлое
	context.SetCookie("refresh_token", "", -1, "/", "", false, true)

	// Отправляем успешный ответ
	context.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func setTokenToCookie(с *gin.Context, token string) {
	с.SetCookie("refresh_token", token, 24*60*60, "/", "", false, true)
}
