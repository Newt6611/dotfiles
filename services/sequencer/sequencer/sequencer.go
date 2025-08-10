// Package sequencer
package sequencer

import (
	"github.com/redis/go-redis/v9"
)

type Sequencer struct {
	redisClient *redis.Client
}

func New(redis *redis.Client) *Sequencer {
	return &Sequencer{
		redisClient: redis,
	}
}
