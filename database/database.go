package database

import (
	models "todo_adv_app/database/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB, err = gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

func ConnectAndCreatre() {
	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(&models.Links{})
	DB.AutoMigrate(&models.Notes{})
}
