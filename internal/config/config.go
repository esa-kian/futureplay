package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Matchmaking struct {
		CompetitionSize int `yaml:"competition_size"`
		WaitTimeSeconds int `yaml:"wait_time_seconds"`
		LevelRange      int `yaml:"level_range"`
	} `yaml:"matchmaking"`
}

func LoadConfig(filePath string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	log.Printf("Loaded configuration: %+v\n", config)
	return config, nil
}
