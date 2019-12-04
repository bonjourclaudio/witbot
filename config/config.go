package config

import "github.com/spf13/viper"

type Config struct {
	WIT_AI_TOKEN	string
	Addr			string
	Port			string
}

func GetConfig() (*Config, error) {

	var config Config

	viper.SetConfigFile(`./config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		return &config, err
	}

	// Fill Config
	config.WIT_AI_TOKEN = viper.GetString("WIT_AI_TOKEN")
	config.Addr = viper.GetString("addr")
	config.Port = viper.GetString("port")

	return &config, nil
}