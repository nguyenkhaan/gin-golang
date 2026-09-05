package db

import (
	"cloudian/cloudian-restful/internal/config"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB 

func ConnectDB() {
	var err error 
	dsn := config.GetEnvVar("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn) , &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database") 
	}
}