package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yourusername/yourprojectname/models"
	"github.com/yourusername/yourprojectname/repositories"
	"github.com/yourusername/yourprojectname/utils"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		userRepo: repositories.NewUserRepository(db),
	}
}

func (us *UserService) CreateUser(c *gin.Context, username, password string) (*models.User, error) {
	if username == "" || password == "" {
		return nil, errors.New("username and password cannot be empty")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: username,
		Password: hashedPassword,
	}

	err = us.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (us *UserService) GetUserByID(id uint) (*models.User, error) {
	return us.userRepo.FindByID(id)
}

func (us *UserService) GetUserByUsername(username string) (*models.User, error) {
	return us.userRepo.FindByUsername(username)
}

func (us *UserService) AuthenticateUser(username, password string) (*models.User, error) {
	user, err := us.userRepo.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("incorrect password")
	}

	return user, nil
}