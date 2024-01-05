package grpc

import (
	"google.golang.org/grpc"
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
