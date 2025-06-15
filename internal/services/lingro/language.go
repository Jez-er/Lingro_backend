package lingro_services

import "Astral-Back-End/internal/repositories"

type LingroLanguageService struct {
	LingroVocabularyRepository *repositories.LingroLanguagesRepository
}

func NewLingroLanguageService(vocabRepo *repositories.LingroLanguagesRepository) *LingroLanguageService {
	return &LingroLanguageService{
		LingroVocabularyRepository: vocabRepo,
	}
}
func (vs *LingroLanguageService) GetById(languageID uint) (*repositories.LingroLanguages, error) {
	lang, err := vs.LingroVocabularyRepository.GetLanguageById(languageID)
	if err != nil {
		return nil, err
	}
	return &lang, nil
}