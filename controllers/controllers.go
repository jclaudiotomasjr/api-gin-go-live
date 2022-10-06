package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jclaudiotomasjr/api-gin-go-live/api-go-gin-go-live-quiz/database"
	"github.com/jclaudiotomasjr/api-gin-go-live/api-go-gin-go-live-quiz/models"
)

func CreateScore(c *gin.Context) {
	var input models.Score
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error()})
		return
	}
	database.DB.Create(&input)
	c.JSON(http.StatusOK, gin.H{
		"status": "Score lançado com sucesso!"})
}

func AllScores(c *gin.Context) {
	var scores []models.Score
	database.DB.Find(&scores)
	c.JSON(http.StatusOK, scores)
}

func GetScoreId(c *gin.Context) {

	var score models.Score

	id := c.Params.ByName("id")
	database.DB.First(&score, id)
	if score.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Score não encontrado"})
		return
	}
	c.JSON(http.StatusOK, score)
	
}

func GetScoreCpf(c *gin.Context) {
	var score models.Score

	cpf := c.Param("cpf")
	database.DB.Where(&models.Score{Cpf:cpf}).First(&score)
	if score.Cpf == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status":"CPF não encontrado"})
			return	
		}
	c.JSON(http.StatusOK, score)

}

func DeleteScore(c *gin.Context){
	var score models.Score
	id := c.Params.ByName("id")
	database.DB.First(&score, id)
	if score.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status":"Score não encontrado!"})
	}
	database.DB.Delete(&score, id)
	c.JSON(http.StatusOK, gin.H{"status":"Score deletado com sucesso!"})

}
