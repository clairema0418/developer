package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AudioSegment struct {
	ID        uint `gorm:"primaryKey"`
	StartTime float64
	EndTime   float64
}

type TextSegment struct {
	ID          uint `gorm:"primaryKey"`
	AudioSegmentID uint
	Text        string
}

type VectorSegment struct {
	ID          uint `gorm:"primaryKey"`
	TextSegmentID uint
	Vector      []float64 `gorm:"type:float[]"`
}

var dbConnection *gorm.DB

func init() {
	dsn := "host=localhost user=postgres password=postgres dbname=audio_processing port=5432 sslmode=disable TimeZone=UTC"
	var err error
	dbConnection, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = dbConnection.AutoMigrate(&AudioSegment{}, &TextSegment{}, &VectorSegment{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}

func storeInDatabase(audioSegment AudioSegment, textSegment TextSegment, vectorSegment VectorSegment) error {
	tx := dbConnection.Begin()

	err := tx.Create(&audioSegment).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to store audio segment: %v", err)
	}

	textSegment.AudioSegmentID = audioSegment.ID
	err = tx.Create(&textSegment).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to store text segment: %v", err)
	}

	vectorSegment.TextSegmentID = textSegment.ID
	err = tx.Create(&vectorSegment).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to store vector segment: %v", err)
	}

	tx.Commit()
	return nil
}