package main

import (
	"log"
	"os"

	"github.com/avanti-dvp/ms-saudacoes-aleatorias/database"
	"github.com/avanti-dvp/ms-saudacoes-aleatorias/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializa a conexão com o banco de dados
	database.ConnectDatabase()

	// Cria um router Gin com as configurações padrão
	router := gin.Default()

	// CORS:
	// ⚠️ Não pode usar AllowCredentials=true junto com AllowOrigins="*".
	// Como a API normalmente usa Authorization: Bearer (sem cookies), deixamos AllowCredentials=false.
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"}, // Permite todas as origens
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: false,
	}))

	// Define as rotas da API
	api := router.Group("/api")
	{
		// Rota para cadastrar um novo cumprimento
		// Ex: POST /api/saudacoes
		api.POST("/saudacoes", handlers.CreateGreeting)

		// Rota para obter um cumprimento aleatório
		// Ex: GET /api/saudacoes/aleatorio
		api.GET("/saudacoes/aleatorio", handlers.GetRandomGreeting)
	}

	// Porta dinâmica (bom para deploy em plataformas como Koyeb):
	// Se PORT não existir, usa 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Inicia o servidor
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}