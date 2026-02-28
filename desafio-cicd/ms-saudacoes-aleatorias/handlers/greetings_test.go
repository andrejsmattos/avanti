package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/avanti-dvp/ms-saudacoes-aleatorias/database"
	"github.com/avanti-dvp/ms-saudacoes-aleatorias/models"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// setupTestDB inicializa um banco de dados em memória para os testes.
func setupTestDB(t *testing.T) {
	// Usando o driver do SQLite em memória
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Falha ao conectar ao banco de dados em memória: %v", err)
	}

	// Migra o schema do banco
	err = db.AutoMigrate(&models.Greeting{})
	if err != nil {
		t.Fatalf("Falha ao migrar o banco de dados: %v", err)
	}

	// Disponibiliza a conexão do banco de dados para os handlers
	database.DB = db

	// Insere um dado de teste
	initialGreeting := models.Greeting{Text: "Olá, Mundo dos Testes!"}
	if err := database.DB.Create(&initialGreeting).Error; err != nil {
		t.Fatalf("Falha ao inserir dado de teste: %v", err)
	}
}

// setupRouter configura o roteador do Gin para os testes.
func setupRouter() *gin.Engine {
	// Desabilita o modo de debug do Gin para um log mais limpo nos testes
	gin.SetMode(gin.TestMode)

	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/saudacoes", CreateGreeting)
		api.GET("/saudacoes/aleatorio", GetRandomGreeting)
	}
	return router
}

func TestCreateGreeting(t *testing.T) {
	// Configura o ambiente de teste
	setupTestDB(t)
	router := setupRouter()

	// Cria o corpo da requisição
	newGreeting := `{"text": "Olá, Mundo dos Testes!"}`
	req, _ := http.NewRequest(http.MethodPost, "/api/saudacoes", bytes.NewBufferString(newGreeting))
	req.Header.Set("Content-Type", "application/json")

	// Cria um ResponseRecorder para capturar a resposta
	w := httptest.NewRecorder()

	// Executa a requisição
	router.ServeHTTP(w, req)

	// Verifica o código de status
	if w.Code != http.StatusCreated {
		t.Errorf("Esperava o código de status %d, mas obteve %d", http.StatusCreated, w.Code)
	}

	// Verifica o corpo da resposta
	var response map[string]models.Greeting
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Falha ao decodificar a resposta JSON: %v", err)
	}

	if data, ok := response["data"]; !ok || data.Text != "Olá, Mundo dos Testes!" {
		t.Errorf("Corpo da resposta inesperado. Obtido: %s", w.Body.String())
	}
}

func TestGetRandomGreeting(t *testing.T) {
	// Configura o ambiente de teste
	setupTestDB(t)
	router := setupRouter()

	// Cria a requisição
	req, _ := http.NewRequest(http.MethodGet, "/api/saudacoes/aleatorio", nil)

	// Cria um ResponseRecorder
	w := httptest.NewRecorder()

	// Executa a requisição
	router.ServeHTTP(w, req)

	// Verifica o código de status
	if w.Code != http.StatusOK {
		t.Errorf("Esperava o código de status %d, mas obteve %d", http.StatusOK, w.Code)
	}

	// Verifica o corpo da resposta
	var response map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("Falha ao decodificar a resposta JSON: %v", err)
	}

	// Como só temos uma saudação no banco de teste, a resposta deve ser ela
	if saudacao, ok := response["saudação"]; !ok || saudacao != "Olá, Mundo dos Testes!" {
		t.Errorf("Esperava a saudação 'Olá, Mundo dos Testes!', mas obteve '%s'", saudacao)
	}
}
