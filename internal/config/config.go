package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Project struct {
	Name   string `yaml:"name"`
	Author string `yaml:"author"`
}

type Grpc struct {
	Address string `yaml:"address"`
}

type Database struct {
	Driver   string `yaml:"driver"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SslMode  string `yaml:"sslmode"`
}

type Metrics struct {
	Address string `yaml:"address"`
	Pattern string `yaml:"pattern"`
}

type Kafka struct {
	Brokers []string `yaml:"brokers"`
	Topic   string   `yaml:"topic"`
}

type Config struct {
	Project  Project  `yaml:"project"`
	Grpc     Grpc     `yaml:"grpc"`
	Database Database `yaml:"database"`
	Metrics  Metrics  `yaml:"metrics"`
	Kafka    Kafka    `yaml:"kafka"`
}

func Read(path string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
