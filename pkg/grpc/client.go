package grpc

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type socketClient struct {
	connection *grpc.ClientConn
}

type SocketClient interface {
	GetConnection() *grpc.ClientConn
}

type SocketClientConfig struct {
	Target string
}

func ProvideGRPCSocketClient(config SocketClientConfig) (SocketClient, error) {
	grpcConnection, err := grpc.Dial(
		config.Target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &socketClient{
		connection: grpcConnection,
	}, nil
}

func (s socketClient) GetConnection() *grpc.ClientConn {
	return s.connection
}
