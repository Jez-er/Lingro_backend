package lingro_controllers

import (
	lingro_services "Astral-Back-End/internal/services/lingro"
	"Astral-Back-End/internal/utils"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary      Add a new word to vocabulary
// @Description  Adds a new word with translations to a specific vocabulary
// @Tags         lingro, words
// @Accept       json
// @Produce      json
// @Param        word  body  lingro_services.NewWordRequest  true  "Word Data"
// @Success      200  {object}  map[string]string  "Word added"
// @Failure      400  {object}  map[string]string  "Invalid request"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /lingro/words/ [post]
func NewWord(context *gin.Context) {
	var newWord lingro_services.NewWordRequest
	db := utils.ORM(context)
	VocabularyService := lingro_services.NewLingroWordsService(db.LingroWordsRepo)

	if err := context.ShouldBindJSON(&newWord); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := VocabularyService.NewWord(newWord); err != nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "Word added"})
}

// @Summary      Get words by vocabulary ID
// @Description  Retrieves all words for a specific vocabulary, with optional field filtering
// @Tags         lingro, words
// @Accept       json
// @Produce      json
// @Param        user_id  path  int  true  "Vocabulary ID"
// @Param        fields   query string false "Comma-separated list of fields to include (id,word,translate,variant,scores)"
// @Success      200  {array}  map[string]interface{}  "List of words"
// @Failure      400  {object}  map[string]string  "Invalid user ID"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /lingro/words/{user_id} [get]
func GetWordsByVocabularyID(context *gin.Context) {
		userID := context.Param("user_id")

    fieldsParam := context.DefaultQuery("fields", "id,word,translate,variant,scores")
    fields := make(map[string]bool)
    
    for _, field := range strings.Split(fieldsParam, ",") {
        fields[strings.TrimSpace(field)] = true
    }

    db := utils.ORM(context)
    VocabularyService := lingro_services.NewLingroWordsService(db.LingroWordsRepo)

    userIDUint, err := strconv.ParseUint(userID, 10, 32)
    if err != nil {
        print(err.Error())
        context.JSON(400, gin.H{"error": "Invalid user ID"})
        return
    }
    words, err := VocabularyService.GetByVocabularyId(uint(userIDUint))
    if err != nil {
        context.JSON(500, gin.H{"error": err.Error()})
        return
    }

    filteredWords := make([]map[string]interface{}, 0, len(words))
    
    for _, word := range words {
        filteredWord := make(map[string]interface{})
        
        if fields["id"] {
            filteredWord["id"] = word.Id
        }
        if fields["word"] {
            filteredWord["word"] = word.Word
        }
        if fields["translate"] {
            filteredWord["translate"] = word.Translate
        }
        if fields["variant"] {
            filteredWord["variant"] = word.Variant
        }
        if fields["scores"] {
            filteredWord["scores"] = word.Scores
        }
      
        if len(filteredWord) > 0 {
            filteredWords = append(filteredWords, filteredWord)
        }
    }

    context.JSON(200, filteredWords)
}

// @Summary      Get words by vocabulary ID
// @Description  Retrieves all words for a specific vocabulary, with optional field filtering
// @Tags         lingro, words
// @Accept       json
// @Produce      json
// @Param        user_id  path  int  true  "Vocabulary ID"
// @Param        fields   query string false "Comma-separated list of fields to include (id,word,translate,variant,scores)"
// @Success      200  {array}  map[string]interface{}  "List of words"
// @Failure      400  {object}  map[string]string  "Invalid user ID"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /lingro/words/{user_id} [get]
func GetWordsByID(context *gin.Context) {
		userID := context.Param("word_id")

    fieldsParam := context.DefaultQuery("fields", "id,word,translate,variant,scores")
    fields := make(map[string]bool)
    
    for _, field := range strings.Split(fieldsParam, ",") {
        fields[strings.TrimSpace(field)] = true
    }

    db := utils.ORM(context)
    VocabularyService := lingro_services.NewLingroWordsService(db.LingroWordsRepo)

    userIDUint, err := strconv.ParseUint(userID, 10, 32)
    if err != nil {
        print(err.Error())
        context.JSON(400, gin.H{"error": "Invalid word ID"})
        return
    }
    words, err := VocabularyService.GetById(uint(userIDUint))
    if err != nil {
        context.JSON(500, gin.H{"error": err.Error()})
        return
    }

	context.JSON(200, gin.H{
		"id":        words.Id,
		"word":      words.Word,
		"translate": words.Translate,
		"variant":   words.Variant,
		"scores":    words.Scores,
	})
}

// @Summary      Update a word
// @Description  Updates a word and its translations by word ID
// @Tags         lingro, words
// @Accept       json
// @Produce      json
// @Param        word_id  path  int  true  "Word ID"
// @Param        word  body  lingro_services.EditWordRequest  true  "Updated Word Data"
// @Success      200  {object}  map[string]string  "Word updated"
// @Failure      400  {object}  map[string]string  "Invalid word ID or request"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /lingro/words/{word_id} [put]
func UpdateWord(context *gin.Context) {
	var updateWord lingro_services.EditWordRequest
	wordID := context.Param("word_id")

	db := utils.ORM(context)
	VocabularyService := lingro_services.NewLingroWordsService(db.LingroWordsRepo)

	wordIDUint, err := strconv.ParseUint(wordID, 10, 32)
	if err!= nil {
		print(err.Error())
		context.JSON(400, gin.H{"error": "Invalid word ID"})
		return
	}

    if err := context.ShouldBindJSON(&updateWord); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := VocabularyService.EditWord(uint(wordIDUint), updateWord); err!= nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "Word updated"})
}

// @Summary      Set scores for a word
// @Description  Updates the scores for a word and adjusts its variant accordingly
// @Tags         lingro, words
// @Accept       json
// @Produce      json
// @Param        word_id  path  int  true  "Word ID"
// @Param        scores   path  int  true  "Scores value"
// @Success      200  {object}  map[string]string  "Scores updated"
// @Failure      400  {object}  map[string]string  "Invalid word ID or scores"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /lingro/words/scores/{word_id}/{scores}/ [put]
func SetScores(context *gin.Context) {
	wordID := context.Param("word_id")
	scores := context.Param("scores")

	db := utils.ORM(context)
	VocabularyService := lingro_services.NewLingroWordsService(db.LingroWordsRepo)

	wordIDUint, err := strconv.ParseUint(wordID, 10, 32)
	if err!= nil {
		print(err.Error())
		context.JSON(400, gin.H{"error": "Invalid word ID"})
	}

	scoresInt, err := strconv.Atoi(scores)
	if err!= nil {
		print(err.Error())
		context.JSON(400, gin.H{"error": "Invalid scores"})
	}

	if err := VocabularyService.SetScores(uint(wordIDUint), scoresInt); err!= nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return  
	}

	context.JSON(200, gin.H{"message": "Scores updated"})
}

// @Summary      Delete a word
// @Description  Deletes a word by its ID
// @Tags         lingro, words
// @Accept       json
// @Produce      json
// @Param        word_id  path  int  true  "Word ID"
// @Success      200  {object}  map[string]string  "Word deleted"
// @Failure      400  {object}  map[string]string  "Invalid word ID"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /lingro/words/{word_id} [delete]
func DeleteWord(context *gin.Context) {
	wordID := context.Param("word_id")

	db := utils.ORM(context)
	VocabularyService := lingro_services.NewLingroWordsService(db.LingroWordsRepo)

	wordIDUint, err := strconv.ParseUint(wordID, 10, 32)
	if err!= nil {
		print(err.Error())
		context.JSON(400, gin.H{"error": "Invalid word ID"})
		return	
	}

	if err := VocabularyService.DeleteWord(uint(wordIDUint)); err!= nil {
		context.JSON(500, gin.H{"error": err.Error()})
		return
	}

	context.JSON(200, gin.H{"message": "Word deleted"})
}
