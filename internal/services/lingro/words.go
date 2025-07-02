package lingro_services

import "Astral-Back-End/internal/repositories"

type LingroWordsService struct {
	VocabularyRepository        *repositories.VocabularyRepository
}

func NewLingroWordsService(vocabRepo *repositories.VocabularyRepository) *LingroWordsService {
	return &LingroWordsService{
		VocabularyRepository:        vocabRepo,
	}
}

type NewWordRequest struct {
	VocabID    uint     `json:"vocabulary_id" binding:"required"`
	Word      string   `json:"word" binding:"required"`
	Translate []string `json:"translate" binding:"required"`
}

type EditWordRequest struct {
	Word      string   `json:"word" binding:"required"`
	Translate []string `json:"translate" binding:"required"`
}

func (vs *LingroWordsService) NewWord(word NewWordRequest) error {
	return vs.VocabularyRepository.Create(word.VocabID, word.Word, word.Translate)
}

func (vs *LingroWordsService) GetByVocabularyId(vocabID uint) ([]repositories.LingroWords, error) {
	return vs.VocabularyRepository.GetWordsByVocabularyID(vocabID)
}

func (vs *LingroWordsService) GetById(ID uint) (repositories.LingroWords, error) {
	wordPtr, err := vs.VocabularyRepository.GetById(ID)
	if err != nil {
		return repositories.LingroWords{}, err
	}
	if wordPtr == nil {
		return repositories.LingroWords{}, nil
	}
	return *wordPtr, nil
}

func (vs *LingroWordsService) EditWord(id uint, word EditWordRequest) error {
	return vs.VocabularyRepository.EditWord(id, word.Word, word.Translate)
}

func (vs *LingroWordsService) SetScores(wordId uint, scores int) error {
	var variant string
	if scores >= 150 {
		variant = "master"
	} else if scores >= 50 {
		variant = "learning"
	} else {
		variant = "new"
	}
	err := vs.VocabularyRepository.SetScores(wordId, scores)
	if err != nil {
		return err
	}
	err = vs.VocabularyRepository.UpdateVariant(wordId, repositories.VariantType(variant))
	if err != nil {
		return err
	}
	return nil
}

func (vs *LingroWordsService) DeleteWord(wordID uint) error {
	return vs.VocabularyRepository.DeleteWord(wordID)
}

