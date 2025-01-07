package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Token    string   `yaml:"token"`
	Guilds   []string `yaml:"guilds"`
	Channels []string `yaml:"channels"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

func (c *Config) GetToken() string {
	return c.Token
}

func (c *Config) GetGuilds() []string {
	return c.Guilds
}

func (c *Config) GetChannels() []string {
	return c.Channels
}
