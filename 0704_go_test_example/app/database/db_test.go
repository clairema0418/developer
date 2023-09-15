```go
package database_test

import (
	"testing"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/stretchr/testify/assert"
	"app/database"
)

func TestDatabaseConnection(t *testing.T) {
	db, err := gorm.Open("postgres", "host=localhost user=gorm dbname=gorm password=gorm sslmode=disable")
	assert.NoError(t, err)

	err = database.DB.DB().Ping()
	assert.NoError(t, err)
}
```