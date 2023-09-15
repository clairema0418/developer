package message_handler

import (
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"generate_questions_api/database"
	"generate_questions_api/python_generator"
)

const questionGenerationCompleted = "question_generation_completed"

type QuestionGenerationPayload struct {
	QuestionCount int
}

func SetupMessageHandler() {
	pubSub := gochannel.NewGoChannel(gochannel.Config{}, middleware.Retry{})
	router, err := message.NewRouter(message.RouterConfig{}, pubSub)
	if err != nil {
		panic(err)
	}

	router.AddNoPublisherHandler(
		"question_generation_handler",
		questionGenerationCompleted,
		pubSub,
		func(msg *message.Message) ([]*message.Message, error) {
			var payload QuestionGenerationPayload
			err := json.Unmarshal(msg.Payload, &payload)
			if err != nil {
				return nil, err
			}

			questions, err := python_generator.CallPythonScript(payload.QuestionCount)
			if err != nil {
				return nil, err
			}

			err = database.SaveQuestionsToDatabase(questions)
			if err != nil {
				return nil, err
			}

			err = database.UpdateProgressInDB(payload.QuestionCount)
			if err != nil {
				return nil, err
			}

			return nil, nil
		},
	)

	go func() {
		if err := router.Run(); err != nil {
			panic(err)
		}
	}()
}