package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/yourusername/yourprojectname/models"
)

type AccountRepository interface {
	Create(account *models.Account) error
	GetByID(id uint) (*models.Account, error)
	Update(account *models.Account) error
	Delete(id uint) error
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) Create(account *models.Account) error {
	return r.db.Create(account).Error
}

func (r *accountRepository) GetByID(id uint) (*models.Account, error) {
	account := &models.Account{}
	err := r.db.First(account, id).Error
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (r *accountRepository) Update(account *models.Account) error {
	return r.db.Save(account).Error
}

func (r *accountRepository) Delete(id uint) error {
	return r.db.Delete(&models.Account{}, id).Error
}