```go
package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var DB *gorm.DB

func InitDB() {
	var err error
	pgUser := os.Getenv("POSTGRES_USER")
	pgPassword := os.Getenv("POSTGRES_PASSWORD")
	pgDB := os.Getenv("POSTGRES_DB")
	pgHost := os.Getenv("POSTGRES_HOST")
	pgPort := os.Getenv("POSTGRES_PORT")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", pgHost, pgUser, pgDB, pgPassword, pgPort)

	DB, err = gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	fmt.Println("Successfully connected to database.")
}
```