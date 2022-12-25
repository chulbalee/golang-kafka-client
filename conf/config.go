package conf

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Type     string `yaml:"type"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DB       string `yaml:"db"`
	} `yaml:"database"`

	KafkaClient struct {
		BootstrapServers []string `yaml:"bootstrap-servers"`
		GroupId          string   `yaml:"group-id"`
		Topics           []string `yaml:"topics"`
	} `yaml:"kafka-client"`
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
