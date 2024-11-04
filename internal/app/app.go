package app

import (
	"context"
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

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}
	app.initDeps(ctx)

	return app, nil
}

func (a *App) Run() error {
	return a.runGrpcServer()
}

func (a *App) initDeps(ctx context.Context) error {
	var deps [3]func(context.Context) error = [3]func(context.Context) error{
		a.initConfig,
		a.initGrpcServer,
		a.initServiceProvider,
	}

	for _, f := range deps {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if err := config.Load(".env"); err != nil {
		return err
	}
	return nil
}

func (a *App) initGrpcServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer()

	reflection.Register(a.grpcServer)
	topic.RegisterTopicV1Server(a.grpcServer, a.serviceProvider.TopicServer())
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) runGrpcServer() error {
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
