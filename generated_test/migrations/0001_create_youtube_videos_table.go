package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yourusername/yourprojectname/utils/db_connector"
)

func main() {
	db := db_connector.Connect()
	defer db.Close()

	err := createYouTubeVideosTable(db)
	if err != nil {
		fmt.Println("Error creating YouTubeVideos table:", err)
	} else {
		fmt.Println("YouTubeVideos table created successfully")
	}
}

func createYouTubeVideosTable(db *gorm.DB) error {
	type YouTubeVideo struct {
		ID            uint   `gorm:"primary_key"`
		URL           string `gorm:"type:text;not null"`
		AudioFilePath string `gorm:"type:text"`
		Transcription string `gorm:"type:text"`
	}

	return db.CreateTable(&YouTubeVideo{}).Error
}