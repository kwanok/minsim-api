package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Database Database `yaml:"database"`
}

type Database struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Port     string `yaml:"port"`
}

func NewConfig(cfgFile string) *Config {
	if cfgFile == "" {
		log.Fatalln("Error reading config file")
	}

	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	config := &Config{}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	log.Println("config loaded")
	return config
}
