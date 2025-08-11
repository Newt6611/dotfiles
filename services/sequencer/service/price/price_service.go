package price

import (
	"github.com/charmbracelet/log"
	"github.com/strike-finance/strike-v2-backend/pkg/models"
	"github.com/strike-finance/strike-v2-backend/pkg/redis"
	"github.com/strike-finance/strike-v2-backend/services/price"
	"github.com/strike-finance/strike-v2-backend/services/price/binance"
	"github.com/strike-finance/strike-v2-backend/services/price/bybit"
	"github.com/strike-finance/strike-v2-backend/services/price/gate"
	"github.com/strike-finance/strike-v2-backend/services/price/mexc"
	"github.com/strike-finance/strike-v2-backend/services/price/okx"
	"github.com/strike-finance/strike-v2-backend/services/sequencer/config"
)

type PriceService struct {
	redisClient *redis.Redis
	exchanges   []price.Exchange
	doneCh      chan struct{}
}

func NewService(config *config.Config, redis *redis.Redis) *PriceService {

	cfg := parseConfig(config)
	exchanges := []price.Exchange{
		binance.New(&cfg),
		okx.New(&cfg),
		bybit.New(&cfg),
		mexc.New(&cfg),
		gate.New(&cfg),
	}

	return &PriceService{
		redisClient: redis,
		exchanges:   exchanges,
		doneCh:      make(chan struct{}),
	}
}

func (p *PriceService) Active() {
	// Create priceFeed channel for receiving price feed from all exchanges
	priceFeedCh := make(chan price.PriceFeed, 1)

	// Active all exchanges
	for _, exchange := range p.exchanges {
		go exchange.Active(priceFeedCh)
	}

	// Shutdown all exchanges
	defer func() {
		for _, exchange := range p.exchanges {
			exchange.Shutdown()
		}
	}()

	var lastTimeUpdated int64 = 0
	for {
		select {
		case <-p.doneCh:
			log.Info("Price Feed Loop Stopped")
			return

		case priceFeed := <-priceFeedCh:
			if priceFeed.TimeInMilli <= lastTimeUpdated {
				continue
			}

			log.Infof("Exchange: %s Pair: %s Price: %f", priceFeed.Exchange, priceFeed.Pair, priceFeed.Price)

			// Push to stream
			p.redisClient.PushPriceEvent(models.PriceEvent{
				Exchange:  priceFeed.Exchange,
				Pair:      priceFeed.Pair,
				Price:     priceFeed.Price,
				Timestamp: priceFeed.TimeInMilli,
			})
		}
	}
}

func (p *PriceService) Shutdown() {
	select {
	case <-p.doneCh:
	default:
		close(p.doneCh)
	}
}
