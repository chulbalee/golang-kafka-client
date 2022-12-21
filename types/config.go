package types

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	Type     string `yaml:"type"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type KafkaClientConfig struct {
	BootstrapServers []string `yaml:"bootstrap-servers"`
	GroupId          string   `yaml:"group-id"`
	Topics           string   `yaml:"topics"`
}

type Config struct {
	DbConfig          DatabaseConfig    `yaml:"database"`
	KafkaClientConfig KafkaClientConfig `yaml:"kafka-client"`
}

func (config *Config) LoadConfig(SERVER_CONFIG_PATH string) *Config {
	configFile, err := ioutil.ReadFile(SERVER_CONFIG_PATH)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, config)
	if err != nil {
		panic(err)
	}

	return config
}
