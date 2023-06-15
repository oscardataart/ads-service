package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	LoadConfig("config_test.yaml")

	expected := Config{
		Server: struct {
			Port   string `yaml:"port"`
			ApiKey string `yaml:"api-key"`
		}{
			Port:   "8080",
			ApiKey: "abc123",
		},
		Redis: struct {
			Address string `yaml:"address"`
		}{
			Address: "localhost:6379",
		},
		Ads: struct {
			URL string `yaml:"url"`
		}{
			URL: "https://www.example.com/sm_ads.json",
		},
	}

	assert.Equal(t, expected, ServiceConfig)
}
