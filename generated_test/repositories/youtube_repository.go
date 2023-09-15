package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/yourusername/yourprojectname/models"
)

type YouTubeRepository interface {
	Create(video *models.YouTubeVideo) error
	FindByID(id string) (*models.YouTubeVideo, error)
	Update(video *models.YouTubeVideo) error
}

type youtubeRepository struct {
	db *gorm.DB
}

func NewYouTubeRepository(db *gorm.DB) YouTubeRepository {
	return &youtubeRepository{db: db}
}

func (r *youtubeRepository) Create(video *models.YouTubeVideo) error {
	return r.db.Create(video).Error
}

func (r *youtubeRepository) FindByID(id string) (*models.YouTubeVideo, error) {
	var video models.YouTubeVideo
	err := r.db.Where("id = ?", id).First(&video).Error
	if err != nil {
		return nil, err
	}
	return &video, nil
}

func (r *youtubeRepository) Update(video *models.YouTubeVideo) error {
	return r.db.Save(video).Error
}