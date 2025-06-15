package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secret = os.Getenv("JWT_SECRET")

func GetSecret() string {
	return secret
}

const (
	AccessTokenDuration  = 15 * time.Minute
	RefreshTokenDuration = 7 * 24 * time.Hour
)

type Claims struct {
	UserID    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	UserEmail string `json:"user_email"`
	jwt.RegisteredClaims
}

func GenerateTokens(userID int, username, useremail string) (string, string, error) {
	accessClaims := Claims{
		UserID:    userID,
		UserName:  username,
		UserEmail: useremail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "AstralEcosystem",
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	refreshClaims := Claims{
		UserID:    userID,
		UserEmail: useremail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "AstralEcosystem",
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(secret))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetSecret()), nil // Используем GetSecret вместо аргумента
	})

	if err != nil {
		fmt.Printf("Error parsing token: %v\n", err)
		return nil, err
	}
	if !token.Valid {
		fmt.Println("Token is invalid")
		return nil, fmt.Errorf("invalid token")
	}

	fmt.Printf("Claims: %+v\n", claims) // Более информативный вывод
	return claims, nil
}
