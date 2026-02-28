package main

import (
	"log"

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

	// A configuração abaixo permite todas as origens (AllowAllOrigins).
	// É uma configuração liberal, ideal para APIs públicas.
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permite todas as origens
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	// Uma forma ainda mais simples para permitir TODAS as origens é:
	// router.Use(cors.Default()) // <- Alternativa simples que já permite '*'

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

	// Inicia o servidor na porta 8080
	// Você pode acessar em http://localhost:8080
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
