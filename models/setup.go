package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=localhost user=postgres dbname=postgres port=5432 sslmode=disable timezone=Europe/Amsterdam"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect to database!")
	}

	database.AutoMigrate(&Entrada{})

	DB = database
}
