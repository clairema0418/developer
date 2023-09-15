package repositories

import (
	"github.com/jinzhu/gorm"
	"models"
)

type DBReaderRepository struct {
	db *gorm.DB
}

func NewDBReaderRepository(db *gorm.DB) *DBReaderRepository {
	return &DBReaderRepository{db: db}
}

func (r *DBReaderRepository) GetAccount(accountID uint) (*models.Account, error) {
	var account models.Account
	err := r.db.Where("id = ?", accountID).First(&account).Error
	if err != nil {
		return nil, err
	}
	return &account, nil
}

func (r *DBReaderRepository) ListAccounts() ([]*models.Account, error) {
	var accounts []*models.Account
	err := r.db.Find(&accounts).Error
	if err != nil {
		return nil, err
	}
	return accounts, nil
}