package app

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	AppName  string `mapstructure:"app_name"`
	Port     string
	Database DatabaseConfig `mapstructure:",squash"`
}

type DatabaseConfig struct {
	Name     string `mapstructure:"database_name"`
	Username string `mapstructure:"database_username"`
	Password string `mapstructure:"database_password"`
	Host     string `mapstructure:"database_host"`
}

func LoadConfiguration() (Config, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading configuration file : %s\n", err)
		return Config{}, nil
	}

	config := Config{}
	viper.Unmarshal(&config)
	return config, nil
}
