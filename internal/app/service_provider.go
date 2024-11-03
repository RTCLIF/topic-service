package app

// Реализация Dependency Injection

import (
	"log"

	"github.com/RTCLIF/topic-service/internal/api/topic"
	"github.com/RTCLIF/topic-service/internal/config"
	"github.com/RTCLIF/topic-service/internal/service"
	topicService "github.com/RTCLIF/topic-service/internal/service/topic"
)

type serviceProvider struct {
	grpcConfig    config.GRPCConfig
	topiceService service.TopicService
	topicServer   *topic.TopicServerImpl
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GrpcConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGrpcConfig()
		if err != nil {
			log.Fatalf("failed to initialize grpcConfig: %s", err)
		}

		s.grpcConfig = cfg
	}
	return s.grpcConfig
}

func (s *serviceProvider) TopicService() service.TopicService {
	if s.topiceService == nil {
		s.topiceService = topicService.NewTopicServiceImpl()
	}

	return s.topiceService
}

func (s *serviceProvider) TopicServer() *topic.TopicServerImpl {
	if s.topicServer == nil {
		s.topicServer = topic.NewTopicServer(
			s.TopicService(),
		)
	}

	return s.topicServer
}
