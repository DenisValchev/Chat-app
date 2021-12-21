package utils

import "github.com/spf13/viper"

type Config struct {
	DBName     string
	DBUser     string
	DBPassword string
	Port       string
}

var config *Config

func getConfig() *Config {
	return config
}

func LoadConfigVars() (*Config, error) {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		return new(Config), err
	}

	config = &Config{
		DBName:     viper.GetString("DB_NAME"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		Port:       viper.GetString("PORT"),
	}
	return config, nil
}
