package config

import (
	"github.com/spf13/viper"
)

func LoadConfig() (*New, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config *New
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}
