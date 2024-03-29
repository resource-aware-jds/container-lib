package grpc

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	connection *grpc.ClientConn
}

type Client interface {
	GetConnection() *grpc.ClientConn
}

type ClientConfig struct {
	Target string
}

func ProvideClient(config ClientConfig) (Client, error) {
	logrus.Info("Debug Config - ", config.Target)
	grpcConnection, err := grpc.Dial(
		config.Target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &client{
		connection: grpcConnection,
	}, nil
}

func (s client) GetConnection() *grpc.ClientConn {
	return s.connection
}
