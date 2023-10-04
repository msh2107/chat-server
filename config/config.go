package config

import "github.com/ilyakaznacheev/cleanenv"

// Config - is the config for the app
type Config struct {
	GRPC `yaml:"grpc"`
}

// GRPC - is the config for gRPC
type GRPC struct {
	Host     string `yaml:"host" env-default:"localhost"`
	GRPCPort string `yaml:"grpc_port" env-default:"50051"`
}

// NewConfig - returns new config for the app
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, err
	}

	return cfg, err
}
