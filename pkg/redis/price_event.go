package redis

import (
	"context"
	"encoding/json"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"
	"github.com/strike-finance/strike-v2-backend/pkg/consts"
	"github.com/strike-finance/strike-v2-backend/pkg/models"
)

func (s *Redis) PushPriceEvent(priceEvent models.PriceEvent) error {
	jsonValue, err := json.Marshal(priceEvent)
	if err != nil {
		return err
	}
	return s.client.XAdd(context.Background(), &redis.XAddArgs{
		ID:     "*", // Let redis generate the ID
		Stream: consts.PriceQueueEvent,
		Values: map[string]any{
			"data": jsonValue,
		},
	}).Err()
}

func (s *Redis) PullPriceEvent(fn func(models.PriceEvent), doneCh chan struct{}) {
	for {
		select {
		case <-doneCh:
			log.Info("Pull price event done")
			return
		default:
		}

		streams, err := s.client.XReadGroup(context.Background(), &redis.XReadGroupArgs{
			Group:    consts.PriceQueueGruop,
			Consumer: "PriceFeed Consumer",
			Streams:  []string{consts.PriceQueueEvent, ">"},
			Block:    0,
			Count:    100,
		}).Result()
		if err != nil {
			panic(err)
		}

		var messageIDs []string
		for _, stream := range streams {
			for _, msg := range stream.Messages {

				var priceEvent models.PriceEvent
				dataStr := msg.Values["data"].(string)
				err := json.Unmarshal([]byte(dataStr), &priceEvent)
				if err != nil {
					log.Error("Failed to unmarshal price event:", err)
					continue
				}
				fn(priceEvent)
				messageIDs = append(messageIDs, msg.ID)
			}
		}

		// Delete all messages in the batch with a single command
		if len(messageIDs) > 0 {
			if err := s.client.XAck(context.Background(), consts.PriceQueueEvent, consts.PriceQueueGruop, messageIDs...).Err(); err != nil {
				log.Error("Failed to ack messages:", err)
			}

			// Delete all messages in the batch with a single command
			if _, err := s.client.XDel(context.Background(), consts.PriceQueueEvent, messageIDs...).Result(); err != nil {
				log.Error("Failed to delete messages:", err)
			}
		}
	}
}
