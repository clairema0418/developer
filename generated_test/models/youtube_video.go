package models

import (
	"github.com/jinzhu/gorm"
)

type YouTubeVideo struct {
	ID             uint   `gorm:"primary_key"`
	URL            string `gorm:"type:text;not null"`
	AudioFilePath  string `gorm:"type:text"`
	Transcription  string `gorm:"type:text"`
}

func (YouTubeVideo) TableName() string {
	return "youtube_videos"
}

func MigrateYouTubeVideo(db *gorm.DB) {
	db.AutoMigrate(&YouTubeVideo{})
}