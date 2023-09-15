```go
package account_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jinzhu/gorm"
	"app/account"
)

var (
	db *gorm.DB
	repo account.Repository
)

func setup() {
	db, _ = gorm.Open("postgres", "host=localhost user=gorm dbname=gorm sslmode=disable password=gorm")
	repo = account.NewRepository(db)
}

func TestCreateAccount(t *testing.T) {
	setup()
	account := account.Account{Name: "Test", Email: "test@test.com"}
	err := repo.CreateAccount(&account)
	assert.Nil(t, err)
}

func TestGetAccount(t *testing.T) {
	setup()
	account, err := repo.GetAccount(1)
	assert.Nil(t, err)
	assert.NotNil(t, account)
}

func TestUpdateAccount(t *testing.T) {
	setup()
	account := account.Account{Name: "Test Updated", Email: "testupdated@test.com"}
	err := repo.UpdateAccount(1, &account)
	assert.Nil(t, err)
}

func TestDeleteAccount(t *testing.T) {
	setup()
	err := repo.DeleteAccount(1)
	assert.Nil(t, err)
}
```