#!/bin/bash

TABLE_NAME=$1

if [ -z "$TABLE_NAME" ]; then
  echo "Error: Table name is required."
  exit 1
fi

cat <<EOF > "models/${TABLE_NAME}.go"
package models

import (
	"time"
)

type ${TABLE_NAME^} struct {
	ID        uint      \`gorm:"primary_key"\`
	CreatedAt time.Time \`gorm:"not null"\`
	UpdatedAt time.Time \`gorm:"not null"\`
}
EOF

cat <<EOF > "controllers/${TABLE_NAME}_controller.go"
package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourproject/models"
	"github.com/yourusername/yourproject/services"
)

func Create${TABLE_NAME^}(c *gin.Context) {
	var ${TABLE_NAME} models.${TABLE_NAME^}
	if err := c.ShouldBindJSON(&${TABLE_NAME}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.Create${TABLE_NAME^}(&${TABLE_NAME}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, ${TABLE_NAME})
}

func Get${TABLE_NAME^}(c *gin.Context) {
	id := c.Param("id")
	${TABLE_NAME}, err := services.Get${TABLE_NAME^}(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ${TABLE_NAME})
}

func Update${TABLE_NAME^}(c *gin.Context) {
	id := c.Param("id")
	var ${TABLE_NAME} models.${TABLE_NAME^}
	if err := c.ShouldBindJSON(&${TABLE_NAME}); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.Update${TABLE_NAME^}(id, &${TABLE_NAME}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ${TABLE_NAME})
}

func Delete${TABLE_NAME^}(c *gin.Context) {
	id := c.Param("id")
	if err := services.Delete${TABLE_NAME^}(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
EOF

cat <<EOF >> "routes/routes.go"
api.POST("/${TABLE_NAME}s", controllers.Create${TABLE_NAME^})
api.GET("/${TABLE_NAME}s/:id", controllers.Get${TABLE_NAME^})
api.PUT("/${TABLE_NAME}s/:id", controllers.Update${TABLE_NAME^})
api.DELETE("/${TABLE_NAME}s/:id", controllers.Delete${TABLE_NAME^})
EOF

echo "CRUD API for ${TABLE_NAME} has been generated."