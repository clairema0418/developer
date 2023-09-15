package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/yourusername/yourprojectname/models"
)

type TranscriptionRepository interface {
	Create(transcription *models.Transcription) error
	GetByVideoID(videoID uint) (*models.Transcription, error)
	Update(transcription *models.Transcription) error
}

type transcriptionRepository struct {
	db *gorm.DB
}

func NewTranscriptionRepository(db *gorm.DB) TranscriptionRepository {
	return &transcriptionRepository{db: db}
}

func (r *transcriptionRepository) Create(transcription *models.Transcription) error {
	return r.db.Create(transcription).Error
}

func (r *transcriptionRepository) GetByVideoID(videoID uint) (*models.Transcription, error) {
	var transcription models.Transcription
	err := r.db.Where("video_id = ?", videoID).First(&transcription).Error
	if err != nil {
		return nil, err
	}
	return &transcription, nil
}

func (r *transcriptionRepository) Update(transcription *models.Transcription) error {
	return r.db.Save(transcription).Error
}