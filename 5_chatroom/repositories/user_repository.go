package repositories

import (
	"github.com/jinzhu/gorm"
	"chatroom/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{db: database}
}

func (repo *UserRepository) CreateUser(user *models.User) error {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := repo.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) UpdateUser(user *models.User) error {
	return repo.db.Save(user).Error
}

func (repo *UserRepository) DeleteUser(id uint) error {
	return repo.db.Delete(&models.User{}, id).Error
}