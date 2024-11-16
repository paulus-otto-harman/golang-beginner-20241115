package app

import "github.com/spf13/viper"

type Config struct {
	AppName  string
	Port     string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Name     string
	Username string
	Password string
	Host     string
}

func LoadConfiguration() (Config, error) {
	viper.AutomaticEnv()

	return Config{
		AppName: viper.GetString("APP_NAME"),
		Port:    viper.GetString("PORT"),
		Database: DatabaseConfig{
			Name:     viper.GetString("DATABES_NAME"),
			Username: viper.GetString("DATABASE_USERNAME"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			Host:     viper.GetString("DATABASE_HOST"),
		},
	}, nil
}
