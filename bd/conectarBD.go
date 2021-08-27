package bd

import (
	"fmt"

	"github.com/emiliano080591/go-dragon-ball/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "emiliano:God080591.@tcp(127.0.0.1:3306)/test_dragon?charset=utf8mb4&parseTime=True&loc=Local"

func InitialMigration() {
	GetConnection()
	DB.AutoMigrate(&models.Personaje{})
}

func GetConnection() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
}
