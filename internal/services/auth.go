package services

import (
	"Astral-Back-End/internal/repositories"
	bcrypt "Astral-Back-End/pkg/crypt"
	"Astral-Back-End/pkg/jwt"
	"fmt"
)

type AuthService struct {
	UserService *UserService
}

func NewAuthService(userService *UserService) *AuthService {
	return &AuthService{userService}
}

func (as *AuthService) Registrarion(user *repositories.User) error {
	userData := *user
	userData.Password = bcrypt.HashPassword(user.Password)

	checkUser, err := as.UserService.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}
	if checkUser != nil {
		return fmt.Errorf("User with this email already exists")
	}

	return as.UserService.CreateUser(&userData)
}

func (as *AuthService) Login(email, password string) (string, string, *repositories.User, error) {
	user, err := as.UserService.GetUserByEmail(email)
	if err != nil {
		return "", "", nil, fmt.Errorf("user with this email does not exist")
	}

	if !bcrypt.ComparePassword(user.Password, password) {
		return "", "", nil, fmt.Errorf("invalid password")
	}

	accessToken, refreshToken, _ := jwt.GenerateTokens(int(user.Id), user.Name, user.Email)

	return accessToken, refreshToken, user, nil
}

func (as *AuthService) Refresh(refreshToken string) (string, string, *repositories.User, error) {
	// Валидируем refresh-токен
	claims, err := jwt.ValidateToken(refreshToken)
	if err != nil {
		return "", "", nil, fmt.Errorf("invalid or expired refresh token")
	}

	// Получаем пользователя по ID из claims
	user, err := as.UserService.GetUserByEmail(claims.UserEmail)
	if err != nil {
		return "", "", nil, fmt.Errorf("user not found")
	}

	// Генерируем новые токены
	newAccessToken, newRefreshToken, err := jwt.GenerateTokens(claims.UserID, user.Name, user.Email)
	if err != nil {
		return "", "", nil, fmt.Errorf("failed to generate tokens: %v", err)
	}

	return newAccessToken, newRefreshToken, user, nil
}
