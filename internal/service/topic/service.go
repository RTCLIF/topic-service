package topic

// здесь можно объявить интерфейс репозитория, а не в модуле репозитория

type TopicServiceImpl struct {
}

func NewTopicServiceImpl() *TopicServiceImpl {
	return &TopicServiceImpl{}
}
