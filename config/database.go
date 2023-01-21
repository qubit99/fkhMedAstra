package config

import (
	"gorm.io/gorm"
	"log"
	"os"
)
import "gorm.io/driver/sqlite"

func DatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("medastra.db"), &gorm.Config{})
	if err != nil {
		log.Println("error creating db", err)
		os.Exit(1)
	}
	return db
}
