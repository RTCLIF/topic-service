package topic

import "github.com/RTCLIF/topic-service/internal/service"

// здесь можно объявить интерфейс репозитория, а не в модуле репозитория

var _ service.TopicService = (*topicServiceImpl)(nil) // эта штука показывает, реализует ли наша структура нужный интерфейс

type topicServiceImpl struct {
}

func NewTopicServiceImpl() *topicServiceImpl {
	return &topicServiceImpl{}
}
