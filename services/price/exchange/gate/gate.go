// Package gate
package gate

import (
	"context"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gateio/gateapi-go/v6"
	"github.com/govalues/decimal"
	"github.com/strike-finance/strike-v2-backend/services/price/config"
	"github.com/strike-finance/strike-v2-backend/services/price/exchange"
)

type Gate struct {
	client     *gateapi.APIClient
	config     *config.Config
	pairs      []string
	doneCh     chan struct{}
	pairsCache map[string]bool // Optimize filter speed when fetch prices back from OKX
}

func New(config *config.Config) *Gate {
	formattedPairs := formatPairs(config.Pairs)
	pairsCache := make(map[string]bool)
	for _, pair := range formattedPairs {
		pairsCache[pair] = true
	}
	return &Gate{
		client:     gateapi.NewAPIClient(gateapi.NewConfiguration()),
		config:     config,
		pairs:      formattedPairs,
		pairsCache: pairsCache,
		doneCh:     make(chan struct{}),
	}
}

func (b *Gate) GetName() string {
	return "Gate"
}

func (b *Gate) Active(priceFeedCh chan<- exchange.PriceFeed) {
	ticker := time.NewTicker(exchange.PriceFeedInterval)
	for {
		select {
		case <-b.doneCh:
			log.Info("Gate Active stopped")
			return
		default:
		}

		symbols, _ , err := b.client.SpotApi.ListTickers(context.Background(), nil)
		if err != nil {
			log.Error(err)
			continue
		}

		for _, symbol := range symbols{
			// Filter out pairs not in config
			if !b.pairsCache[symbol.CurrencyPair] {
				continue
			}

			price, _ := decimal.Parse(symbol.Last)
			priceFeedCh <- exchange.PriceFeed{
				Exchange:    b.GetName(),
				Pair:        symbol.CurrencyPair,
				TimeInMilli: time.Now().UnixMilli(),
				Price:       price,
			}
		}

		<-ticker.C
	}
}

func (b *Gate) Shutdown() {
	select {
	case <-b.doneCh:
	default:
		close(b.doneCh)
	}
}
