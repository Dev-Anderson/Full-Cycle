package models

import "gorm.io/gorm"

type Carro struct {
	gorm.Model
	Modelo string `json:"modelo"`
	Marca  string `json:"marca"`
	Cor    string `json:"cor"`
}
