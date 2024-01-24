package db

import (
	"log"

	"admin_svc/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	log.Println("DB initialized")
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Payment{})
	db.AutoMigrate(&models.Category{})

	return Handler{db}
}
