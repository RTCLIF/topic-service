package topic

import (
	"github.com/RTCLIF/topic-service/internal/service"
	pb "github.com/RTCLIF/topic-service/pkg/topic_v1"
)

// здесь можно объявить интерфейс сервиса

type TopicServerImpl struct {
	pb.UnimplementedTopicV1Server
	topicService service.TopicService
}

func NewTopicServer(topicService service.TopicService) *TopicServerImpl {
	return &TopicServerImpl{
		topicService: topicService,
	}
}
