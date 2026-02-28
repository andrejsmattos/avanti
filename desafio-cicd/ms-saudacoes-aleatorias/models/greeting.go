package models

import "gorm.io/gorm"

// Greeting representa a estrutura de um cumprimento no banco de dados.
type Greeting struct {
	gorm.Model        // Inclui campos como ID, CreatedAt, UpdatedAt, DeletedAt
	Text       string `json:"text" binding:"required"` // O texto do cumprimento
}
