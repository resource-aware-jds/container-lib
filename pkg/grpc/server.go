package grpc

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"os"
)

type socketServer struct {
	listener   net.Listener
	grpcServer *grpc.Server
}

type SocketServer interface {
	Serve()
	GetGRPCServer() *grpc.Server
}

type ServerConfig struct {
	UnixSocketPath string
}

func ProvideGRPCSocketServer(c ServerConfig) (SocketServer, func(), error) {
	listener, err := net.Listen("unix", c.UnixSocketPath)
	if err != nil {
		return nil, nil, err
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.UnaryInterceptor(grpcUnaryInterceptor),
	)

	cleanup := func() {
		grpcServer.GracefulStop()
		os.Remove(c.UnixSocketPath)
	}

	return &socketServer{
		listener:   listener,
		grpcServer: grpcServer,
	}, cleanup, nil
}

func (s socketServer) Serve() {
	go func() {
		logrus.Info("GRPC Server is Listening on: ", s.listener.Addr())
		s.grpcServer.Serve(s.listener)
	}()
}

func (s socketServer) GetGRPCServer() *grpc.Server {
	return s.grpcServer
}
