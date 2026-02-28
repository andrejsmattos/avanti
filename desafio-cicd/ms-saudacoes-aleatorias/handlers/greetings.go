package handlers

import (
	"net/http"

	"github.com/avanti-dvp/ms-saudacoes-aleatorias/database"
	"github.com/avanti-dvp/ms-saudacoes-aleatorias/models"

	"github.com/gin-gonic/gin"
)

// CreateGreeting cadastra um novo cumprimento no banco de dados.
func CreateGreeting(c *gin.Context) {
	var input models.Greeting

	// Faz o bind do JSON recebido para a struct Greeting
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Cria o novo cumprimento no banco de dados
	greeting := models.Greeting{Text: input.Text}
	if err := database.DB.Create(&greeting).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Falha ao cadastrar a saudação."})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": greeting})
}

// GetRandomGreeting retorna um cumprimento aleatório do banco de dados.
func GetRandomGreeting(c *gin.Context) {
	var greeting models.Greeting

	// GORM com SQLite usa a função RANDOM() para ordenação aleatória.
	// Para outros bancos como PostgreSQL use `RANDOM()` e MySQL use `RAND()`.
	if err := database.DB.Order("RANDOM()").First(&greeting).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Nenhuma saudação encontrada."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"saudação": greeting.Text})
}
