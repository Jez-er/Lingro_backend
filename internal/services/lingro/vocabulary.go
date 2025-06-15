package lingro_services

import (
	"Astral-Back-End/internal/repositories"
	"fmt"
)

type LingroVocabularyService struct {
	VocabularyRepository *repositories.LingroVocabularyRepository
}

func NewLingroVocabularyService(vocabRepo *repositories.LingroVocabularyRepository) *LingroVocabularyService {
	return &LingroVocabularyService{
		VocabularyRepository: vocabRepo,
	}
}

type CreateVocabularyRequest struct {
	UserID    uint     `json:"user_id" binding:"required"`
	LanguageID uint     `json:"language_id" binding:"required"`
}

func (ls *LingroVocabularyService) Create(request *CreateVocabularyRequest) error {

	data, error := ls.VocabularyRepository.GetByLanguageID(request.LanguageID)
	if error != nil {
		return fmt.Errorf("Vocabulary not created")
	}

	if len(data) != 0  {
		return fmt.Errorf("Vocabulary already exists")
	}

	err := ls.VocabularyRepository.Create(request.UserID, request.LanguageID)

	if err != nil  {
		return fmt.Errorf("Vocabulary not created")
		
	}

	return nil
}

func (ls *LingroVocabularyService) GetByUserID(userID uint) ([]repositories.LingroVocabulary, error) {
	vocab, err := ls.VocabularyRepository.GetByUserID(userID)
	return vocab, err
}

func (ls *LingroVocabularyService) Delete(vocabID uint) error {
	err := ls.VocabularyRepository.Delete(vocabID)
	return err
}