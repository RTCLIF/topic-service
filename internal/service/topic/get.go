package topic

import (
	"context"
	"time"

	"github.com/Bulat147/rtclif-topic-service/internal/model"
)

func (t *TopicServiceImpl) GetTopic(ctx context.Context, id string) (model.Topic, error) {
	return model.Topic{
		Id: id,
		Info: model.TopicInfo{
			Title:       "Title",
			Description: "Desc",
		},
		Status:    "Open",
		CreatedAt: time.Now(),
		UpdatedAt: nil,
	}, nil
}
