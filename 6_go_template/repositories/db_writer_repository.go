package repositories

import (
	"github.com/jinzhu/gorm"
	"models"
)

type DBWriterRepository struct {
	DB *gorm.DB
}

func NewDBWriterRepository(db *gorm.DB) *DBWriterRepository {
	return &DBWriterRepository{DB: db}
}

func (repo *DBWriterRepository) CreateAccount(account *models.Account) error {
	return repo.DB.Create(account).Error
}

func (repo *DBWriterRepository) UpdateAccount(account *models.Account) error {
	return repo.DB.Save(account).Error
}

func (repo *DBWriterRepository) DeleteAccount(accountID uint) error {
	return repo.DB.Delete(&models.Account{}, accountID).Error
}