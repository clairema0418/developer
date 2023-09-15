package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Transcription struct {
	ID               uint `gorm:"primary_key"`
	YouTubeVideoID   uint
	YouTubeVideo     YouTubeVideo
	Transcription    string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (Transcription) TableName() string {
	return "transcriptions"
}

func (t *Transcription) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedAt", time.Now())
	return nil
}

func (t *Transcription) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("UpdatedAt", time.Now())
	return nil
}