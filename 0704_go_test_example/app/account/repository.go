```go
package account

import (
	"github.com/jinzhu/gorm"
	"github.com/myapp/app/models"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) CreateAccount(account *models.Account) error {
	return r.DB.Create(account).Error
}

func (r *Repository) GetAccount(id uint) (*models.Account, error) {
	account := &models.Account{}
	if err := r.DB.First(account, id).Error; err != nil {
		return nil, err
	}
	return account, nil
}

func (r *Repository) UpdateAccount(account *models.Account) error {
	return r.DB.Save(account).Error
}

func (r *Repository) DeleteAccount(id uint) error {
	return r.DB.Delete(&models.Account{}, id).Error
}
```