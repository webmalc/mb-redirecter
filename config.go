package main

import (
	"github.com/spf13/viper"
)

// Config is the logger configuration struct.
type Config struct {
	Port int
}

// NewConfig returns the configuration object.
func NewConfig() *Config {
	return &Config{
		Port: viper.GetInt("port"),
	}
}
