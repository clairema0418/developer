package database

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	ID        uint           `gorm:"primaryKey"`
	Content   string         `gorm:"type:text"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}