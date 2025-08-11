// Package binance
package binance

import (
	"context"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/charmbracelet/log"
	"github.com/govalues/decimal"
	"github.com/strike-finance/strike-v2-backend/services/price"
)

type Binance struct {
	client             *binance.Client
	config             *price.Config
	pairs              []string
	doneCh             chan struct{}
	currentEndpointIdx int
}

func New(config *price.Config) *Binance {
	return &Binance{
		client:             binance.NewClient("", ""),
		config:             config,
		pairs:              formatPairs(config.Pairs),
		doneCh:             make(chan struct{}),
		currentEndpointIdx: 0,
	}
}

func (b *Binance) GetName() string {
	return "Binance"
}

func (b *Binance) Active(priceFeedCh chan<- price.PriceFeed) {
	ticker := time.NewTicker(price.PriceFeedInterval)
	for {
		select {
		case <-b.doneCh:
			log.Info("Binance Active stopped")
			return
		default:
		}

		endpoint := b.getEndpoint()
		b.client.SetApiEndpoint(endpoint)

		symbols, err := b.client.NewListPricesService().Symbols(b.pairs).Do(context.Background())
		if err != nil {
			log.Error(err)
			continue
		}

		for _, symbol := range symbols {
			newPrice, _ := decimal.Parse(symbol.Price)
			priceFeedCh <- price.PriceFeed{
				Exchange:    b.GetName(),
				Pair:        symbol.Symbol,
				TimeInMilli: time.Now().UnixMilli(),
				Price:       newPrice,
			}
		}

		<-ticker.C
	}
}

func (b *Binance) Shutdown() {
	select {
	case <-b.doneCh:
	default:
		close(b.doneCh)
	}
}
