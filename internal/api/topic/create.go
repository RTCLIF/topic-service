package topic

import (
	"context"

	"github.com/RTCLIF/topic-service/internal/converter"

	pb "github.com/RTCLIF/topic-service/pkg/topic_v1"
)

func (i *TopicServerImpl) CreateTopic(ctx context.Context, rq *pb.CreateTopicRq) (*pb.CreateTopicRs, error) {
	id, err := i.topicService.CreateTopic(ctx, *converter.ToTopicInfoFromPb(rq.GetInfo()))
	if err != nil {
		return nil, err
	}
	return &pb.CreateTopicRs{Id: id}, nil
}
