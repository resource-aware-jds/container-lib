package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"os"
)

type Config struct {
	WorkerNodeUnixSocketPath string `envconfig:"WORKER_NODE_UNIX_SOCKET_PATH"`
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
		return nil, err
	}

	var config Config
	envconfig.MustProcess("RAJDS", &config)

	return &config, nil
}
