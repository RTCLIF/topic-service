package converter

// репо <-> сервис конвертеры сюда не вынес, т.к. есть коллизия имен,
// а иначе придется удлиннять имена (указывать не toTopic..., а toRepoTopic..., toServiceTopic... и т.д.)

import (
	"github.com/RTCLIF/topic-service/internal/model"
	pb "github.com/RTCLIF/topic-service/pkg/topic_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToTopicInfoFromPb(topicInfo *pb.TopicInfo) *model.TopicInfo {
	return &model.TopicInfo{
		Title:       topicInfo.Title,
		Description: topicInfo.Description,
	}
}

func ToTopicFromService(topic model.Topic) *pb.Topic {
	var updatedAt *timestamppb.Timestamp // nil - если не указан
	if topic.UpdatedAt != nil {
		updatedAt = timestamppb.New(*topic.UpdatedAt)
	}

	return &pb.Topic{
		Id:        topic.Id,
		Info:      ToTopicInfoFromService(topic.Info),
		Status:    topic.Status,
		CreatedAt: timestamppb.New(topic.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToTopicInfoFromService(topicInfo model.TopicInfo) *pb.TopicInfo {
	return &pb.TopicInfo{
		Title:       topicInfo.Title,
		Description: topicInfo.Description,
	}
}
