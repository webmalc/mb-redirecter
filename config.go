package main

import (
	"github.com/spf13/viper"
)

// Config is the logger configuration struct.
type Config struct {
	Port     int
	BaseURL  string
	IsProd   bool
	APIUrl   string
	APIToken string
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	return &Config{
		Port:     viper.GetInt("port"),
		BaseURL:  viper.GetString("base_url"),
		APIUrl:   viper.GetString("api_url"),
		APIToken: viper.GetString("api_token"),
		IsProd:   viper.GetBool("is_prod"),
	}
}
