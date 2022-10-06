package models

import (
	"gorm.io/gorm"
)

type Score struct {
	gorm.Model
	Cpf         string `json:"cpf"`
	Sector      string `json:"sector"`
	Punctuation int64  `json:"punctuation"`
}
