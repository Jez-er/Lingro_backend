package routes

import (
	lingro_controllers "Astral-Back-End/internal/controllers/lingro"

	"github.com/gin-gonic/gin"
)

func LingroRouter(router *gin.RouterGroup) {
	groupe := router.Group("/lingro")
	vocabularyGroupe := groupe.Group("/vocabulary")
	{
		vocabularyGroupe.POST("/", lingro_controllers.NewVocabulary)
		vocabularyGroupe.GET("/:userId", lingro_controllers.GetVocabularyByUserId)
		vocabularyGroupe.DELETE("/:vocabId", lingro_controllers.DeleteVocabulary)
	}
	wordGroupe := groupe.Group("/word")
	{
		wordGroupe.POST("/", lingro_controllers.NewWord)
		wordGroupe.GET("/:user_id", lingro_controllers.GetWordsByVocabularyID)
		wordGroupe.GET("/word/:word_id", lingro_controllers.GetWordsByID)
		wordGroupe.PUT("/:word_id", lingro_controllers.UpdateWord)
		wordGroupe.PUT("/scores/:word_id/:scores", lingro_controllers.SetScores)
		wordGroupe.DELETE("/:word_id", lingro_controllers.DeleteWord)
	}
}
