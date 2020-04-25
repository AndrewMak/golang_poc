package utils

import (
	"log"
	"os"

	"github.com/andrewmak/poc_golang/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func ConnectDb() *gorm.DB {
	err := godotenv.Load()
	err != nil {
		log.Fatalf("error loading .env file")
	}
	dsn := os.Getenv("dsn")
	
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("error connecting to database", err)
		panic(err)
	}
	// defet db.Close()
	db.AutoMigrate(&domain.User{})

	return db
}
