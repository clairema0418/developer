package workers

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"youtube_audio_to_text_converter/models"
	"youtube_audio_to_text_converter/pubsub"
	"youtube_audio_to_text_converter/services"
	"youtube_audio_to_text_converter/utils"
)

type YouTubeWorker struct {
	db               *gorm.DB
	youtubeService   *services.YouTubeService
	whisperService   *services.WhisperService
	subscriber       *gochannel.Subscriber
	publisher        *gochannel.Publisher
}

func NewYouTubeWorker(db *gorm.DB, youtubeService *services.YouTubeService, whisperService *services.WhisperService) *YouTubeWorker {
	subscriber := pubsub.NewSubscriber()
	publisher := pubsub.NewPublisher()

	return &YouTubeWorker{
		db:               db,
		youtubeService:   youtubeService,
		whisperService:   whisperService,
		subscriber:       subscriber,
		publisher:        publisher,
	}
}

func (worker *YouTubeWorker) Start() error {
	youtubeVideoCreatedSubscription, err := worker.subscriber.Subscribe(pubsub.YoutubeVideoCreated)
	if err != nil {
		return errors.Wrap(err, "failed to subscribe to youtube_video_created")
	}

	go worker.processYouTubeVideoCreated(youtubeVideoCreatedSubscription)

	return nil
}

func (worker *YouTubeWorker) processYouTubeVideoCreated(messages <-chan *message.Message) {
	for msg := range messages {
		video := &models.YouTubeVideo{}
		if err := msg.Unmarshal(video); err != nil {
			msg.Nack()
			continue
		}

		audioFilePath, err := utils.DownloadYouTubeAudio(video.URL)
		if err != nil {
			msg.Nack()
			continue
		}

		video.AudioFilePath = audioFilePath
		worker.youtubeService.UpdateAudioFilePath(video)

		audioExtractedMsg := message.NewMessage(msg.UUID(), nil)
		if err := audioExtractedMsg.Marshal(video); err != nil {
			msg.Nack()
			continue
		}

		worker.publisher.Publish(pubsub.YoutubeAudioExtracted, audioExtractedMsg)
		msg.Ack()
	}
}