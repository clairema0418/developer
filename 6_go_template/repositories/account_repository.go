package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/yourusername/yourprojectname/models"
)

type AccountRepository interface {
	CreateAccount(account *models.Account) error
	UpdateAccount(account *models.Account) error
	DeleteAccount(accountID uint) error
	GetAccount(accountID uint) (*models.Account, error)
	ListAccounts() ([]*models.Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) CreateAccount(account *models.Account) error {
	return r.db.Create(account).Error
}

func (r *accountRepository) UpdateAccount(account *models.Account) error {
	return r.db.Save(account).Error
}

func (r *accountRepository) DeleteAccount(accountID uint) error {
	return r.db.Delete(&models.Account{ID: accountID}).Error
}

func (r *accountRepository) GetAccount(accountID uint) (*models.Account, error) {
	var account models.Account
	err := r.db.First(&account, accountID).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) ListAccounts() ([]*models.Account, error) {
	var accounts []*models.Account
	err := r.db.Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}