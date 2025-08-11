// Package redis
package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/strike-finance/strike-v2-backend/pkg/consts"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(client *redis.Client) *Redis {
	return &Redis{
		client: client,
	}
}

func (s *Redis) Init() {
	s.client.XGroupCreateMkStream(context.Background(), consts.PriceQueueEvent, consts.PriceQueueGruop, "$").Err()
}
