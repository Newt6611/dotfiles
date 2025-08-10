package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/charmbracelet/log"
	"github.com/strike-finance/strike-v2-backend/pkg/models"
	"github.com/strike-finance/strike-v2-backend/services/price/config"
	"github.com/strike-finance/strike-v2-backend/services/price/exchange"
	"github.com/strike-finance/strike-v2-backend/services/price/exchange/binance"
	"github.com/strike-finance/strike-v2-backend/services/price/redis"
	"github.com/strike-finance/strike-v2-backend/services/sequencer/sequencer"
)

func main() {
	var wg sync.WaitGroup

	log.Info("Service Price Started")

	// Init config
	cfg := config.Init()

	// Init redis client
	redisClient := redis.NewClient(cfg)

	// Sequencer client
	sequencerClient := sequencer.New(redisClient.GetClient())

	// Setup all exchanges
	exchanges := []exchange.Exchange{
		binance.New(cfg),
	}

	// Create priceFeed channel for receiving price feed from all exchanges
	priceFeedCh := make(chan exchange.PriceFeed, 1)

	// Active all exchanges
	ActiveAllExchanges(exchanges, priceFeedCh)

	wg.Add(1)

	doneCh := make(chan struct{})

	// Start price feed loop
	go func() {
		defer wg.Done()

		var lastTimeUpdated int64 = 0
		for {
			select {
			case <-doneCh:
				log.Info("Price feed loop stopped")
				return

			case priceFeed := <-priceFeedCh:
				if priceFeed.TimeInMilli <= lastTimeUpdated {
					continue
				}
				msg := fmt.Sprintf("Exchange: %s Pair: %s Price: %f", priceFeed.Exchange, priceFeed.Pair, priceFeed.Price)
				log.Info(msg)

				// Push price event to sequencer
				PushPriceEvent(sequencerClient, priceFeed)
			}
		}
	}()

	// Handle signal like Ctrl+C
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh

	log.Info("Shutting down service price...")

	ShutdownAllExchanges(exchanges)

	// Waiting for price feed loop closed
	close(doneCh)
	wg.Wait()

	log.Info("Shutdown complete")
}

func ActiveAllExchanges(exchanges []exchange.Exchange, priceFeedCh chan<- exchange.PriceFeed) {
	for _, exchange := range exchanges {
		go exchange.Active(priceFeedCh)
	}
}

func ShutdownAllExchanges(exchanges []exchange.Exchange) {
	for _, exchange := range exchanges {
		exchange.Shutdown()
	}
}

func PushPriceEvent(s *sequencer.Sequencer, priceFeed exchange.PriceFeed) {
	s.PushPriceEvent(models.PriceEvent{
		Exchange:  priceFeed.Exchange,
		Pair:      priceFeed.Pair,
		Price:     priceFeed.Price,
		Timestamp: priceFeed.TimeInMilli,
	})
}
