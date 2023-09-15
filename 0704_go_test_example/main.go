```go
package main

import (
	"github.com/gin-gonic/gin"
	"./app/config"
	"./app/database"
	"./app/account"
)

func main() {
	config.Load()
	database.Connect()

	router := gin.Default()

	account.RegisterRoutes(router)

	router.Run()
}
```