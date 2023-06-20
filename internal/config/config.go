package config

import (
	"fmt"
	"os"

	db "foodApp/pkg/db"
	rb "foodApp/pkg/messageBroker/rabbitMq/config"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		HTTP struct {
			Port int `yaml:"port"`
		} `yaml:"http"`
	} `yaml:"server"`
	DBConfig       *db.Config         `yaml:"database"`
	RabbitMqConfig *rb.RabbitMQConfig `yaml:"rabbitmq"`
}

func Load(configFile string) (*Config, error) {
	appConfig := &Config{}
	if _, err := os.Stat(configFile); err != nil {
		log.Err(fmt.Errorf("could not find local.yaml in directory: %w", err))
	}

	bytes, err := os.ReadFile(configFile)
	if err != nil {
		log.Err(fmt.Errorf("error reading config file: %w", err))
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}
	if err = yaml.Unmarshal(bytes, &appConfig); err != nil {
		log.Err(fmt.Errorf("error unmarshalling config: %w", err))
		return nil, fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return appConfig, nil
}
