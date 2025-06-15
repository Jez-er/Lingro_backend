package repositories

import (
	"time"

	"gorm.io/gorm"
)

type LingroVocabulary struct {
	Id        uint      `gorm:"primaryKey"`
	LanguageID uint      `json:"language_id" binding:"required" gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	User      User      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type LingroVocabularyRepository struct {
	db *gorm.DB
}

func NewLingroVocabularyRepository(db *gorm.DB) *LingroVocabularyRepository {
	return &LingroVocabularyRepository{db}
}

func (vr *LingroVocabularyRepository) Create(userId uint, languageId uint) error {
	vocabulary := LingroVocabulary{
		LanguageID: languageId,
		UserID:    userId,
	}

	return vr.db.Create(&vocabulary).Error
}

func (vr *LingroVocabularyRepository) GetByUserID(userID uint) ([]LingroVocabulary, error) {
	var vocabularies []LingroVocabulary
	err := vr.db.Where("user_id = ?", userID).Find(&vocabularies).Error
	return vocabularies, err
}

func (vr *LingroVocabularyRepository) GetByLanguageID(langID uint) ([]LingroVocabulary, error) {
	var vocabularies []LingroVocabulary
	err := vr.db.Where("language_id = ?", langID).Find(&vocabularies).Error
	return vocabularies, err
}

func (vr *LingroVocabularyRepository) Delete(vocabularyID uint) error {
	return vr.db.Delete(&LingroVocabulary{}, vocabularyID).Error
}
