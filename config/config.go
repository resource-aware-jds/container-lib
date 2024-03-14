package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/resource-aware-jds/container-lib/pkg/grpc"
	"github.com/resource-aware-jds/container-lib/pkg/taskrunner"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
)

type Config struct {
	Debug                      bool   `envconfig:"DEBUG" default:"false"`
	ProfilingToolsListeningURL string `envconfig:"PROFILING_TOOLS_LISTENING_URL" default:"[::1]:31111"`
	ContainerGRPCListeningURL  string `envconfig:"CONTAINER_GRPC_LISTENING_URL" default:":31236"`
	WorkerNodeReceiverGRPCURL  string `envconfig:"WORKER_NODE_RECEIVER_GRPC_URL" default:"host.docker.internal:31237"`
	ImageURL                   string `envconfig:"IMAGE_URL" required:"true"`
	InitialTaskRunner          int    `envconfig:"INITIAL_TASK_RUNNER" default:"1"`
	ContainerId                string
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
	loadContainerId(&config)

	return &config, nil
}

func ProvideGRPCSocketServerConfig(config *Config) grpc.ServerConfig {
	return grpc.ServerConfig{
		GRPCServerListeningAddr: config.ContainerGRPCListeningURL,
	}
}

func ProvideTaskRunnerPoolConfig(config *Config) taskrunner.PoolConfig {
	return taskrunner.PoolConfig{
		NumberOfInitialTaskRunner: config.InitialTaskRunner,
	}
}

func ProvideGRPCSocketClientConfig(config *Config) grpc.ClientConfig {
	return grpc.ClientConfig{
		Target: config.WorkerNodeReceiverGRPCURL,
	}
}

func loadContainerId(config *Config) {
	containerId, err := exec.Command("hostname").Output()
	if err != nil {
		logrus.Warn("[TaskRunner Manager] Fail to retrieve hostname with error (%s)", err.Error())
	}
	config.ContainerId = string(containerId)
}
