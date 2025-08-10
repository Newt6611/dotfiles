// Package redis
package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/strike-finance/strike-v2-backend/pkg/consts"
	"github.com/strike-finance/strike-v2-backend/services/price/config"
)

type RedisClient struct {
	config *config.Config
	client *redis.Client
}

func NewClient(config *config.Config) *RedisClient {
	urlOpt, err := redis.ParseURL(config.RedisURL)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(urlOpt)

	err = client.Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}

	_ = client.XGroupCreateMkStream(
		context.Background(),
		consts.PriceQueueEvent,
		consts.AggregatorQueueGruop,
		"$",
	).Err()

	return &RedisClient{
		config: config,
		client: client,
	}
}

func (r *RedisClient) GetClient() *redis.Client {
	return r.client
}

func (r *RedisClient) Shutdown() {
	r.client.Shutdown(context.Background()).Result()
}
