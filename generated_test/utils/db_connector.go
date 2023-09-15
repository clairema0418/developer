package utils

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
)

func ConnectDB() (*gorm.DB, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath("../config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	dbUser := viper.GetString("database.user")
	dbPassword := viper.GetString("database.password")
	dbName := viper.GetString("database.name")
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetInt("database.port")

	dbURI := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", dbHost, dbPort, dbUser, dbName, dbPassword)

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("Error connecting to the database: %s", err)
		return nil, err
	}

	return db, nil
}