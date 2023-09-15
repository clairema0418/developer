package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/yourusername/yourprojectname/models"
	"github.com/yourusername/yourprojectname/repositories"
	"github.com/yourusername/yourprojectname/events"
	"github.com/yourusername/yourprojectname/publishers"
)

type AccountService struct {
	repo       *repositories.AccountRepository
	publisher  *publishers.AccountPublisher
}

func NewAccountService(db *gorm.DB, publisher *publishers.AccountPublisher) *AccountService {
	return &AccountService{
		repo:       repositories.NewAccountRepository(db),
		publisher:  publisher,
	}
}

func (s *AccountService) CreateAccount(c *gin.Context, account *models.Account) error {
	if err := s.repo.Create(account); err != nil {
		return err
	}

	event := events.NewAccountCreatedEvent(account)
	s.publisher.PublishAccountCreated(c, event)

	return nil
}

func (s *AccountService) GetAccount(c *gin.Context, id uint) (*models.Account, error) {
	return s.repo.GetByID(id)
}

func (s *AccountService) UpdateAccount(c *gin.Context, account *models.Account) error {
	if err := s.repo.Update(account); err != nil {
		return err
	}

	event := events.NewAccountUpdatedEvent(account)
	s.publisher.PublishAccountUpdated(c, event)

	return nil
}

func (s *AccountService) DeleteAccount(c *gin.Context, id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	event := events.NewAccountDeletedEvent(id)
	s.publisher.PublishAccountDeleted(c, event)

	return nil
}