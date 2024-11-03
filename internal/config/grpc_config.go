package config

import (
	"errors"
	"net"
	"os"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type GRPCConfig interface {
	Address() string
}

type grpcConfig struct {
	host string
	port string
}

func NewGrpcConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if host == "" {
		return nil, errors.New("host not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if port == "" {
		return nil, errors.New("port not found")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (c *grpcConfig) Address() string {
	return net.JoinHostPort(c.host, c.port) // host:port
}
