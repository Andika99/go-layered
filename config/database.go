package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"go-layered/models"
)

func Connect(databaseURL string) (db *gorm.DB, err error) {
	dsn := databaseURL
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	database.AutoMigrate(&models.User{})
	
	return database, err
}
