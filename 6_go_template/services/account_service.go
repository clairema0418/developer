package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"

	"models"
	"repositories"
	"messages"
)

type AccountService struct {
	dbWriterRepo repositories.DBWriterRepository
	dbReaderRepo repositories.DBReaderRepository
	publisher    message.Publisher
}

func NewAccountService(db *gorm.DB, pubSub *gochannel.GoChannel) *AccountService {
	accountRepo := repositories.NewAccountRepository(db)
	dbWriterRepo := repositories.NewDBWriterRepository(accountRepo)
	dbReaderRepo := repositories.NewDBReaderRepository(accountRepo)

	publisher := message.NewPublisher(pubSub)

	return &AccountService{
		dbWriterRepo: dbWriterRepo,
		dbReaderRepo: dbReaderRepo,
		publisher:    publisher,
	}
}

func (s *AccountService) CreateAccount(c *gin.Context, account *models.Account) error {
	err := s.dbWriterRepo.CreateAccount(account)
	if err != nil {
		return err
	}

	msg := messages.NewAccountCreated(account.ID)
	err = s.publisher.Publish(messages.AccountCreatedTopic, msg)
	return err
}

func (s *AccountService) UpdateAccount(c *gin.Context, account *models.Account) error {
	err := s.dbWriterRepo.UpdateAccount(account)
	if err != nil {
		return err
	}

	msg := messages.NewAccountUpdated(account.ID)
	err = s.publisher.Publish(messages.AccountUpdatedTopic, msg)
	return err
}

func (s *AccountService) DeleteAccount(c *gin.Context, accountID uint) error {
	err := s.dbWriterRepo.DeleteAccount(accountID)
	if err != nil {
		return err
	}

	msg := messages.NewAccountDeleted(accountID)
	err = s.publisher.Publish(messages.AccountDeletedTopic, msg)
	return err
}

func (s *AccountService) GetAccount(c *gin.Context, accountID uint) (*models.Account, error) {
	return s.dbReaderRepo.GetAccount(accountID)
}

func (s *AccountService) ListAccounts(c *gin.Context) ([]*models.Account, error) {
	return s.dbReaderRepo.ListAccounts()
}