```go
package account_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/jinzhu/gorm"
	"app/account"
)

var (
	service = account.NewService(&gorm.DB{})
)

func TestCreateAccount(t *testing.T) {
	account, err := service.CreateAccount("testUser", "testPassword")
	assert.NoError(t, err)
	assert.NotNil(t, account)
}

func TestGetAccount(t *testing.T) {
	account, err := service.GetAccount(1)
	assert.NoError(t, err)
	assert.NotNil(t, account)
}

func TestUpdateAccount(t *testing.T) {
	account, err := service.UpdateAccount(1, "newTestUser", "newTestPassword")
	assert.NoError(t, err)
	assert.NotNil(t, account)
}

func TestDeleteAccount(t *testing.T) {
	err := service.DeleteAccount(1)
	assert.NoError(t, err)
}
```