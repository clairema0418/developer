package services

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/jinzhu/gorm"
	"youtube_audio_to_text_converter/models"
	"youtube_audio_to_text_converter/repositories"
	"youtube_audio_to_text_converter/utils"
)

type YouTubeService struct {
	db                 *gorm.DB
	youtubeRepository  repositories.YouTubeRepository
	pubsub             *message.Publisher
}

func NewYouTubeService(db *gorm.DB, youtubeRepository repositories.YouTubeRepository, pubsub *message.Publisher) *YouTubeService {
	return &YouTubeService{
		db:                 db,
		youtubeRepository:  youtubeRepository,
		pubsub:             pubsub,
	}
}

func (ys *YouTubeService) CreateYouTubeVideo(url string) (*models.YouTubeVideo, error) {
	video := &models.YouTubeVideo{URL: url}
	err := ys.youtubeRepository.Create(video)
	if err != nil {
		return nil, err
	}

	err = ys.pubsub.Publish("youtube_video_created", message.NewMessage(video.ID, []byte(video.URL)))
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (ys *YouTubeService) ExtractYouTubeAudio(video *models.YouTubeVideo) error {
	audioFilePath, err := utils.DownloadYouTubeAudio(video.URL)
	if err != nil {
		return err
	}

	video.AudioFilePath = audioFilePath
	err = ys.youtubeRepository.Update(video)
	if err != nil {
		return err
	}

	err = ys.pubsub.Publish("youtube_audio_extracted", message.NewMessage(video.ID, []byte(video.AudioFilePath)))
	if err != nil {
		return err
	}

	return nil
}