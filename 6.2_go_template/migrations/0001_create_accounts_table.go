package migrations

import (
	"github.com/jinzhu/gorm"
)

func Migrate0001CreateAccountsTable(db *gorm.DB) {
	db.Exec(`
		CREATE TABLE accounts (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		);
	`)
}