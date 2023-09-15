```go
package account

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Account struct {
	gorm.Model
	ID        uint      `gorm:"primary_key"`
	Name      string    `gorm:"size:255"`
	Email     string    `gorm:"type:varchar(100);unique_index"`
	Password  string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
```