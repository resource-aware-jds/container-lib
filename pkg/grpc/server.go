package grpc

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
)

type server struct {
	listener   net.Listener
	grpcServer *grpc.Server
}

type ServerConfig struct {
	GRPCServerListeningAddr string
}

type Server interface {
	Serve()
	GetGRPCServer() *grpc.Server
}

func ProvideGRPCServer(c ServerConfig) (Server, func(), error) {
	listener, err := net.Listen("tcp", c.GRPCServerListeningAddr)
	if err != nil {
		logrus.Errorf("[GRPC Server] Failed to listen on %s with error %s", c.GRPCServerListeningAddr, err.Error())
		return nil, nil, err
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(grpcUnaryInterceptor),
	)

	cleanup := func() {
		grpcServer.GracefulStop()
	}

	return &server{
		listener:   listener,
		grpcServer: grpcServer,
	}, cleanup, nil
}

func (s server) Serve() {
	go func() {
		logrus.Info("GRPC Server is Listening on: ", s.listener.Addr())
		s.grpcServer.Serve(s.listener)
	}()
}

func (s server) GetGRPCServer() *grpc.Server {
	return s.grpcServer
}
