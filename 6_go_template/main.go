package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"

	"config"
	"controllers"
	"middlewares"
	"models"
	"routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	db, err := config.ConnectDatabase()
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		os.Exit(1)
	}

	migrateDatabase(db)

	router := gin.Default()
	router.Use(middlewares.SetupMiddleware())

	accountController := controllers.NewAccountController(db)
	routes.SetupRoutes(router, accountController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}

func migrateDatabase(db *gorm.DB) {
	err := models.RunMigrations(db)
	if err != nil {
		fmt.Println("Error running migrations:", err)
		os.Exit(1)
	}
}