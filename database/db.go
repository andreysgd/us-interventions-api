package database

import (
	"log"

	"github.com/us-interventions-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func DatabaseConnection()  {
	connectionString := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))
	if err != nil {
		log.Panic("Error trying to connect to database.")
	}
	DB.AutoMigrate(&models.Intervention{})
}