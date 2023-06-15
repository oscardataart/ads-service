package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var (
	ServiceConfig Config
)

type Config struct {
	Server struct {
		Port   string `yaml:"port"`
		ApiKey string `yaml:"api-key"`
	}
	Redis struct {
		Address string `yaml:"address"`
	}
	Ads struct {
		URL string `yaml:"url"`
	}
}

func LoadConfig(configFile string) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	err = yaml.Unmarshal(data, &ServiceConfig)
	if err != nil {
		log.Fatalf("Failed to parse YAML: %v", err)
	}
}
