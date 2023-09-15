package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/yourusername/yourprojectname/utils/db_connector"
)

func main() {
	db, err := db_connector.Connect()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	defer db.Close()

	err = createTranscriptionsTable(db)
	if err != nil {
		panic(fmt.Sprintf("Failed to create transcriptions table: %v", err))
	}
}

func createTranscriptionsTable(db *gorm.DB) error {
	return db.Exec(`
		CREATE TABLE IF NOT EXISTS transcriptions (
			id SERIAL PRIMARY KEY,
			youtube_video_id INTEGER NOT NULL REFERENCES youtube_videos(id),
			text TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		);
	`).Error
}