package initializers

import (
	"errors"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPG_DB() (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	dsn := os.Getenv("PG_DB_URI")
	log.Println("dsn", dsn)

	for i := 0; i < 5; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Connected to PostgreSQL")
			return db, nil
		}
		log.Printf("Failed to connect to PostgreSQL (attempt %d): %v\n", i+1, err)
		time.Sleep(2 * time.Second)
	}

	return nil, errors.New("failed to connect to PostgreSQL after 5 attempts")
}
