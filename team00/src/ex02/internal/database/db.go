package database

import (
	"log"

	"report/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "host=localhost user=name password=pass dbname=anomaly_detection port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	if err = db.AutoMigrate(&models.Anomaly{}); err != nil {
		log.Fatalf("error auto migrating database: %v", err)
	}

	return db
}
