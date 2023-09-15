```go
package account

import (
	"github.com/jinzhu/gorm"
	"app/models"
)

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) CreateAccount(account *models.Account) error {
	return s.DB.Create(account).Error
}

func (s *Service) GetAccount(id uint) (*models.Account, error) {
	var account models.Account
	if err := s.DB.First(&account, id).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (s *Service) UpdateAccount(account *models.Account) error {
	return s.DB.Save(account).Error
}

func (s *Service) DeleteAccount(id uint) error {
	return s.DB.Delete(&models.Account{}, id).Error
}
```