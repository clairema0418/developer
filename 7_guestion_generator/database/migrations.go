package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&Question{})
}