package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/resource-aware-jds/container-lib/pkg/grpc"
	"github.com/resource-aware-jds/container-lib/pkg/taskrunner"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	ContainerUnixSocketPath  string `envconfig:"CONTAINER_UNIX_SOCKET_PATH" default:"/tmp/rajds_container.sock"`
	WorkerNodeUnixSocketPath string `envconfig:"WORKER_NODE_UNIX_SOCKET_PATH" default:"unix:////tmp/rajds_workernode.sock"`
	InitialTaskRunner        int    `envconfig:"INITIAL_TASK_RUNNER" default:"3"`
}

func ProvideConfig() (*Config, error) {
	// Load the .env file only for
	EnvConfigLocation, ok := os.LookupEnv("ENV_CONFIG")
	if !ok {
		EnvConfigLocation = "./.env"
	}

	err := godotenv.Load(EnvConfigLocation)
	if err != nil {
		logrus.Warn("Can't load env file")
	}

	var config Config
	envconfig.MustProcess("RAJDS", &config)

	return &config, nil
}

func ProvideGRPCSocketServerConfig(config *Config) grpc.ServerConfig {
	return grpc.ServerConfig{
		UnixSocketPath: config.ContainerUnixSocketPath,
	}
}

func ProvideTaskRunnerPoolConfig(config *Config) taskrunner.PoolConfig {
	return taskrunner.PoolConfig{
		NumberOfInitialTaskRunner: config.InitialTaskRunner,
	}
}

func ProvideGRPCSocketClientConfig(config *Config) grpc.SocketClientConfig {
	return grpc.SocketClientConfig{
		Target: config.WorkerNodeUnixSocketPath,
	}
}
