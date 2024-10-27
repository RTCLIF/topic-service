package topic

import (
	"context"

	"github.com/Bulat147/rtclif-topic-service/internal/converter"

	pb "github.com/Bulat147/rtclif-topic-service/pkg/topic_v1"
)

func (i *TopicServerImpl) GetTopic(ctx context.Context, rq *pb.GetTopicRq) (*pb.GetTopicRs, error) {
	topic, err := i.topicService.GetTopic(ctx, rq.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.GetTopicRs{Topic: converter.ToTopicFromService(topic)}, nil
}
