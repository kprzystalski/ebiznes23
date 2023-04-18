package database

import (
	"awesomeProject/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("sklep.db"))
	if err != nil {
		panic("Nie bangla")
	}
	db.AutoMigrate(&models.User{})
	return db
}
