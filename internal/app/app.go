package app

import (
	"log"
	"net"

	"github.com/RTCLIF/topic-service/internal/config"
	topic "github.com/RTCLIF/topic-service/pkg/topic_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

func NewApp() (*App, error) {
	app := &App{serviceProvider: newServiceProvider()}
	app.grpcServer = grpc.NewServer()

	reflection.Register(app.grpcServer)
	topic.RegisterTopicV1Server(app.grpcServer, app.serviceProvider.TopicServer())

	if err := config.Load(".env"); err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run() error {
	log.Printf("Grpc Server is running on %s", a.serviceProvider.GrpcConfig().Address())

	listener, err := net.Listen("tcp", a.serviceProvider.GrpcConfig().Address())
	if err != nil {
		return err
	}

	err = a.grpcServer.Serve(listener)
	if err != nil {
		return err
	}

	return nil
}
