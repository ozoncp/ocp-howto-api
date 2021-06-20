package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Project - структура, содержащая параметры проекта
type Project struct {
	Name   string `yaml:"name"`
	Author string `yaml:"author"`
}

// Grpc - структура, содержащая параметры gRPC сервера
type Grpc struct {
	Address string `yaml:"address"`
}

// Database - структура, содержащая параметры базы данных
type Database struct {
	Driver    string `yaml:"driver"`
	Port      string `yaml:"port"`
	Database  string `yaml:"database"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	SslMode   string `yaml:"sslmode"`
	BatchSize int    `yaml:"batch_size"`
}

// Metrics - структура, содержащая параметры метрик
type Metrics struct {
	Address string `yaml:"address"`
	Pattern string `yaml:"pattern"`
}

// Kafka - структура, содержащая параметры для отправки событий в kafka
type Kafka struct {
	Capacity  int      `yaml:"events_queue_capacity"`
	Brokers   []string `yaml:"brokers"`
	Topic     string   `yaml:"topic"`
	Key       string   `yaml:"key"`
	Partition int32    `yaml:"partition"`
}

// Config - структура, содержащая все конфигурируемые параметры
type Config struct {
	Project  Project  `yaml:"project"`
	Grpc     Grpc     `yaml:"grpc"`
	Database Database `yaml:"database"`
	Metrics  Metrics  `yaml:"metrics"`
	Kafka    Kafka    `yaml:"kafka"`
}

// Read считывает конфигурацию из файла
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
