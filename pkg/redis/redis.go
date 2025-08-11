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

func New(client *redis.Client) *Redis {
	r := &Redis{
		client: client,
	}
	r.Init()
	return r
}

func NewFromURL(url string) (*Redis, error) {
	opt, err := redis.ParseURL(url)
	if err != nil {
		return nil, err
	}

	r := &Redis{
		client: redis.NewClient(opt),
	}
	r.Init()
	return r, nil
}

func (s *Redis) Init() {
	s.client.XGroupCreateMkStream(context.Background(), consts.PriceQueueEvent, consts.PriceQueueGruop, "$").Err()
}
