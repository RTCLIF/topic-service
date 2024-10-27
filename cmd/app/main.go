package main

import (
	"log"
	"net"

	"github.com/Bulat147/rtclif-topic-service/internal/api/topic"
	service "github.com/Bulat147/rtclif-topic-service/internal/service/topic"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	dep "github.com/Bulat147/rtclif-topic-service/pkg/topic_v1"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	topicService := service.NewTopicServiceImpl()

	dep.RegisterTopicV1Server(grpcServer, topic.NewTopicServer(topicService))

	log.Println("grpc server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to run server: %s", err)
	}

}
