package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Account{})
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}
	return nil
}