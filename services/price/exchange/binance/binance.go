// Package binance
package binance

import (
	"context"
	"time"

	"github.com/adshao/go-binance/v2"
	"github.com/charmbracelet/log"
	"github.com/govalues/decimal"
	"github.com/strike-finance/strike-v2-backend/services/price/config"
	"github.com/strike-finance/strike-v2-backend/services/price/exchange"
)

type Binance struct {
	client             *binance.Client
	config             *config.Config
	doneCh             chan struct{}
	currentEndpointIdx int
}

func New(config *config.Config) *Binance {
	return &Binance{
		client:             binance.NewClient("", ""),
		config:             config,
		doneCh:             make(chan struct{}),
		currentEndpointIdx: 0,
	}
}

func (b *Binance) GetName() string {
	return "Binance"
}

func (b *Binance) Active(priceFeedCh chan<- exchange.PriceFeed) {
	ticker := time.NewTicker(exchange.PriceFeedInterval)
	for {
		select {
		case <-b.doneCh:
			log.Info("Binance Active stopped")
			return
		default:
		}

		endpoint := b.getEndpoint()
		b.client.SetApiEndpoint(endpoint)

		symbols, err := b.client.NewListPricesService().Symbols(b.config.Pairs).Do(context.Background())
		if err != nil {
			log.Error(err)
		}

		for _, symbol := range symbols {
			price, _ := decimal.Parse(symbol.Price)
			priceFeedCh <- exchange.PriceFeed{
				Exchange:    b.GetName(),
				Pair:        symbol.Symbol,
				TimeInMilli: time.Now().UnixMilli(),
				Price:       price,
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

func (b *Binance) getEndpoint() string {
	b.currentEndpointIdx++
	if b.currentEndpointIdx >= len(b.config.BinanceEndpoints) {
		b.currentEndpointIdx = 0
	}
	return b.config.BinanceEndpoints[b.currentEndpointIdx]
}
