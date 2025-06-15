package repositories

import (
	"gorm.io/gorm"
)

type LingroLanguages struct {
	ID       uint   `gorm:"primaryKey"`
	Language string `json:"language" gorm:"uniqueIndex;not null"` 
	Name     string `json:"name" gorm:"not null"`
}

type LingroLanguagesRepository struct {
	db *gorm.DB
}

func NewLingroLanguagesRepository(db *gorm.DB) *LingroLanguagesRepository {
	return &LingroLanguagesRepository{db: db}
}

func (r *LingroLanguagesRepository) GetAllLanguages() ([]LingroLanguages, error) {
	var languages []LingroLanguages
	result := r.db.Find(&languages)
	return languages, result.Error
}

func (r *LingroLanguagesRepository) GetLanguageById(id uint) (LingroLanguages, error) {
	var language LingroLanguages
	result := r.db.Where("id = ?", id).First(&language)
	return language, result.Error
}

func (r *LingroLanguagesRepository) GetLanguageByCode(code string) (LingroLanguages, error) {
	var language LingroLanguages
	result := r.db.Where("language = ?", code).First(&language)
	return language, result.Error
}