package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	dep "github.com/Bulat147/rtclif-topic-service/pkg/topic_v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MyGrpcServer struct {
	dep.UnimplementedTopicV1Server
}

func (MyGrpcServer) CreateTopic(cnt context.Context, rq *dep.CreateTopicRq) (rs *dep.CreateTopicRs, e error) {
	id, _ := uuid.NewRandom()
	rs = &dep.CreateTopicRs{Id: id.String()}

	return rs, nil
}

func (MyGrpcServer) GetTopic(cnt context.Context, rq *dep.GetTopicRq) (rs *dep.GetTopicRs, e error) {
	id, _ := uuid.NewRandom()
	rs = &dep.GetTopicRs{Topic: &dep.Topic{
		Id:          id.String(),
		Title:       "Random",
		Description: "Random Text",
		Status:      "Open",
		CreatedAt:   timestamppb.Now(),
	}}
	return rs, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen: %s", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	dep.RegisterTopicV1Server(grpcServer, &MyGrpcServer{})

	log.Println("grpc server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to run server: %s", err)
	}

}
