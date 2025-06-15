package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) string {
	password := []byte(pass)

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func ComparePassword(hashedPass, pass string) bool {
	hashedPassword, password := []byte(hashedPass), []byte(pass)
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		return false
	} else {
		return true
	}
}
