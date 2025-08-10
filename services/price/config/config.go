// Package config
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Pairs      []string `mapstructure:"PAIRS"`
	BinanceWss string   `mapstructure:"BINANCE_WSS"`
}

func Init() *Config {
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Read config error : %v", err)
	}
	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	return &cfg
}
