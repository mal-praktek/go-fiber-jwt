package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`
}
