package repositories

import (
	"Astral-Back-End/pkg/database"
)

// All Repositories
type Repositories struct {
	UserRepo              *UserRepository

	// Lingro
	LingroWordsRepo       *VocabularyRepository
	LingroLanguagesRepo         *LingroLanguagesRepository
	LingroVocabularyRepo        *LingroVocabularyRepository
}

func RunDB() *Repositories {
	db := database.InitDB()

	// Migrate the schema
	db.AutoMigrate(&User{}, &LingroWords{}, &LingroLanguages{}, &LingroVocabulary{})

	// Initialize repositories
	userRepo := NewUserRepository(db)

	// Lingro
	vocabularyRepo := NewLingroVocabularyRepository(db) 
	wordRepo := NewVocabularyRepository(db)
	languagesRepo := NewLingroLanguagesRepository(db)

	// Initial data 
	languagesRepo.InitLanguages()

	return &Repositories{
		UserRepo:              userRepo,
		// Lingro
		LingroWordsRepo:       wordRepo,
		LingroLanguagesRepo:         languagesRepo,
		LingroVocabularyRepo:        vocabularyRepo,
	}
}
