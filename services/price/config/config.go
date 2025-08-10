// Package config
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Pairs            []string `mapstructure:"PAIRS"`
	RedisURL         string   `mapstructure:"REDIS_URL"`
	BinanceEndpoints []string `mapstructure:"BINANCE_ENDPOINTS"`
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
