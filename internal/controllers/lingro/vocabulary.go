package lingro_controllers

import (
	lingro_services "Astral-Back-End/internal/services/lingro"
	"Astral-Back-End/internal/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Add a new vocabulary
// @Description  Adds a new vocabulary for a user
// @Tags         lingro, vocabulary
// @Accept       json
// @Produce      json
// @Param        vocabulary  body  lingro_services.CreateVocabularyRequest  true  "Vocabulary Data"
// @Success      200  {object}  map[string]string  "Word added"
// @Failure      400  {object}  map[string]string  "Invalid request"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /lingro/vocabulary/ [post]
func NewVocabulary(ctx *gin.Context) {
	var newVocabulary lingro_services.CreateVocabularyRequest
	db := utils.ORM(ctx)
	VocabularyService := lingro_services.NewLingroVocabularyService(db.LingroVocabularyRepo)

	if err := ctx.ShouldBindJSON(&newVocabulary); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := VocabularyService.Create(&newVocabulary); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "Word added"})
}

// @Summary      Get vocabulary by user ID
// @Description  Retrieves all vocabulary words for a specific user
// @Tags         lingro, vocabulary
// @Accept       json
// @Produce      json
// @Param        userId  path  int  true  "User ID"
// @Success      200  {array}  map[string]interface{}  "List of vocabulary words"
// @Failure      400  {object}  map[string]string  "Invalid user ID"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /lingro/vocabulary/{userId} [get]
func GetVocabularyByUserId(ctx *gin.Context) {
	userId := ctx.Param("userId")
	
	userIDUint, err := strconv.ParseUint(userId, 10, 32)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid user ID"})
		return
	}
	db := utils.ORM(ctx)
	VocabularyService := lingro_services.NewLingroVocabularyService(db.LingroVocabularyRepo)
	LanguageService := lingro_services.NewLingroLanguageService(db.LingroLanguagesRepo)

	words, err := VocabularyService.GetByUserID(uint(userIDUint))
	if err!= nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	filteredWords := make([]map[string]interface{}, 0, len(words))
	for _, word := range words {
		lang, err := LanguageService.GetById(word.LanguageID)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		filteredWord := map[string]interface{}{
			"id":         word.Id,
			"languageId": word.LanguageID,
			"name": lang.Name,
			"code": lang.Language,
		}
		filteredWords = append(filteredWords, filteredWord)
	}

	ctx.JSON(200, filteredWords)
}

// @Summary      Delete a vocabulary word
// @Description  Deletes a vocabulary word by its ID
// @Tags         lingro, vocabulary
// @Accept       json
// @Produce      json
// @Param        vocabId  path  int  true  "Vocabulary ID"
// @Success      200  {object}  map[string]string  "Vocabulary deleted"
// @Failure      400  {object}  map[string]string  "Invalid vocabulary ID"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /lingro/vocabulary/{vocabId} [delete]
func DeleteVocabulary(ctx *gin.Context) {
	vocabId := ctx.Param("vocabId")

	vocabIDUint, err := strconv.ParseUint(vocabId, 10, 32)
	if err!= nil {
		ctx.JSON(400, gin.H{"error": "Invalid vocabulary ID"})
		return
	}
	db := utils.ORM(ctx)
	VocabularyService := lingro_services.NewLingroVocabularyService(db.LingroVocabularyRepo)
	

	if err := VocabularyService.Delete(uint(vocabIDUint)); err!= nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return		
	}

	ctx.JSON(200, gin.H{"message": "Vocabulary deleted"})
}