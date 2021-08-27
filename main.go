package main

import (
	"github.com/emiliano080591/go-dragon-ball/bd"
	"github.com/emiliano080591/go-dragon-ball/handlers"
	_ "gorm.io/gorm"
)

func main() {
	bd.InitialMigration()
	handlers.Manejadores()
}
