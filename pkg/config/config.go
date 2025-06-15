package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func InitConfig() Config {
	file, err := os.Open("config/config.yml")
	if err != nil {
		log.Fatalf("CFG | Could not open config.yml: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("CFG | Unable to decode config.yml: %v", err)
	}
	return config
}
