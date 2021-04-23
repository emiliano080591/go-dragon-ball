package main

import (
	"github.com/emiliano080591/dragon/bd"
	"github.com/emiliano080591/dragon/handlers"
	_ "gorm.io/gorm"
)

func main() {
	bd.InitialMigration()
	handlers.Manejadores()
}
