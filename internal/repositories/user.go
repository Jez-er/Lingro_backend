package repositories

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id                uint               `gorm:"primaryKey"`
	Name              string             `json:"name" binding:"required" gorm:"type:varchar(100)"`
	Email             string             `json:"email" binding:"required,email" gorm:"unique;not null"`
	Password          string             `json:"password" binding:"required" gorm:"not null"`
	IsEmailVerified   bool               `gorm:"default:false"`
	Vocabulary        []LingroVocabulary      `gorm:"foreignKey:UserID"`
	CreatedAt         time.Time          `gorm:"autoCreateTime"`
	UpdatedAt         time.Time          `gorm:"autoUpdateTime"`
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user *User) error {
	return ur.db.Create(user).Error
}

func (ur *UserRepository) GetByEmail(email string) (*User, error) {
	var user User
	err := ur.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
