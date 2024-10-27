package service

import (
	"context"

	"github.com/Bulat147/rtclif-topic-service/internal/model"
)

type TopicService interface {
	CreateTopic(ctx context.Context, topicInfo model.TopicInfo) (string, error)
	GetTopic(ctx context.Context, id string) (model.Topic, error)
}
