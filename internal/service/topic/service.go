package topic

import "github.com/Bulat147/rtclif-topic-service/internal/service"

// здесь можно объявить интерфейс репозитория, а не в модуле репозитория

var _ service.TopicService = (*TopicServiceImpl)(nil) // эта штука показывает, реализует ли наша структура нужный интерфейс

type TopicServiceImpl struct {
}

func NewTopicServiceImpl() *TopicServiceImpl {
	return &TopicServiceImpl{}
}
