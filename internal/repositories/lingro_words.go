package repositories

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

type VariantType string

const (
	NEW      VariantType = "new"
	LEARNING VariantType = "learning"
	MASTER   VariantType = "master"
)

type LingroWords struct {
	Id        uint        `gorm:"primaryKey"`
	Word      string      `json:"word" binding:"required" gorm:"not null"`
	Translate pq.StringArray   `json:"translate" binding:"required" gorm:"type:text[]"`
	Variant   VariantType `json:"variant" binding:"required" gorm:"not null"`
	Scores    int         `gorm:"default:0"`
	VocabularyId    uint        `gorm:"not null"`
	Vocabulary      LingroVocabulary        `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoUpdateTime"`
}

func (v *LingroWords) BeforeCreate(tx *gorm.DB) (err error) {
	if v.Variant == "" {
		v.Variant = NEW
	}
	return nil
}

type VocabularyRepository struct {
	db *gorm.DB
}

func NewVocabularyRepository(db *gorm.DB) *VocabularyRepository {
	return &VocabularyRepository{db}
}

func (vr *VocabularyRepository) Create(vocId uint, word string, translate []string) error {
	vocabulary := LingroWords{
		VocabularyId:    vocId,
		Word:      word,
		Translate: translate,
	}

	return vr.db.Create(&vocabulary).Error
}

func (vr *VocabularyRepository) GetWordsByVocabularyID(vocID uint) ([]LingroWords, error) {
	var vocabularies []LingroWords
	err := vr.db.Where("vocabulary_id = ?", vocID).Find(&vocabularies).Error
	return vocabularies, err
}

func (vr *VocabularyRepository) GetById(wordId uint) (*LingroWords, error) {
	var word LingroWords
	err := vr.db.Where("id = ?", wordId).First(&word).Error
	if err != nil {
		return nil, err
	}
	return &word, nil
}

func (vr *VocabularyRepository) SetScores(wordId uint, scores int) error {
	return vr.db.Model(&LingroWords{}).Where("id =?", wordId).Update("scores", scores).Error
}

func (vr *VocabularyRepository) UpdateVariant(wordId uint, variant VariantType) error {
	return vr.db.Model(&LingroWords{}).Where("id =?", wordId).Update("variant", variant).Error
}

func (vr *VocabularyRepository) EditWord(wordId uint, word string, translate []string) error {
	return vr.db.Model(&LingroWords{}).Where("id =?", wordId).Updates(LingroWords{Word: word, Translate: translate}).Error
}

func (vr *VocabularyRepository ) DeleteWord(wordId uint) error {
	return vr.db.Where("id = ?", wordId).Delete(&LingroWords{}).Error
}