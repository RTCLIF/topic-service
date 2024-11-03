package topic

import (
	"context"
	"time"

	"github.com/RTCLIF/topic-service/internal/model"
)

func (t *topicServiceImpl) GetTopic(ctx context.Context, id string) (model.Topic, error) {
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
