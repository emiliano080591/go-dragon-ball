package models

import "gorm.io/gorm"

type Personaje struct {
	gorm.Model
	Nombre            string `json:"nombre"`
	Primera_aparicion string `json:"primera_aparicion"`
	Descripcion       string `json:"descripcion"`
}
