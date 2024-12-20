package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB     DatabaseConfig
	Server ServerConfig
	Redis  RedisConfig
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type ServerConfig struct {
	Port string
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file: %v", err)
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Printf("Error unmarshalling config: %v", err)
		return nil, err
	}

	return &config, nil
}
