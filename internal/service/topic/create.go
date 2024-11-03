package topic

import (
	"context"

	"github.com/RTCLIF/topic-service/internal/model"
	"github.com/google/uuid"
)

func (t *topicServiceImpl) CreateTopic(ctx context.Context, topicInfo model.TopicInfo) (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
