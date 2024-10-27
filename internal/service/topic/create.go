package topic

import (
	"context"

	"github.com/Bulat147/rtclif-topic-service/internal/model"
	"github.com/google/uuid"
)

func (t *TopicServiceImpl) CreateTopic(ctx context.Context, topicInfo model.TopicInfo) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
