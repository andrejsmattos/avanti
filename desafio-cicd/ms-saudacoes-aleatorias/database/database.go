package database

import (
	"log"

	"github.com/avanti-dvp/ms-saudacoes-aleatorias/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectDatabase inicializa a conexão com o banco de dados SQLite e realiza a migração.
func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("file:greetings.db?_journal_mode=WAL&_mutex=no&_fk=true"), &gorm.Config{})
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados!")
	}

	// AutoMigrate cria a tabela 'greetings' baseada no modelo
	err = db.AutoMigrate(&models.Greeting{})
	if err != nil {
		log.Fatal("Falha ao migrar o banco de dados!")
	}

	DB = db

	// Inicializa a carga inicial de dados
	SeedDatabase()
}

// SeedDatabase insere dados iniciais no banco de dados se ele estiver vazio.
func SeedDatabase() {
	var count int64
	// Conta quantos registros existem na tabela greetings
	DB.Model(&models.Greeting{}).Count(&count)

	// Se o contador for 0, significa que a tabela está vazia
	if count == 0 {
		log.Println("Banco de dados vazio. Inserindo saudações iniciais...")

		// Cria uma lista de saudações
		greetings := []models.Greeting{
			{Text: "Olá"},
			{Text: "Bem-vindo"},
			{Text: "Que a Força esteja com você"},
			{Text: "E aí, tudo certo"},
			{Text: "Live long and prosper"},
			{Text: "Opa, bão"},
			{Text: "Saudações"},
			{Text: "Keep calm and code on"},
			{Text: "Alô, alô"},
		}

		// Insere a lista no banco de dados
		if err := DB.Create(&greetings).Error; err != nil {
			log.Fatalf("Falha ao inserir carga inicial de dados: %v", err)
		}

		log.Println("Carga inicial de saudações inserida com sucesso!")
	} else {
		log.Println("O banco de dados já contém dados. Carga inicial não necessária.")
	}
}
