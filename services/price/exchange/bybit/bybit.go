// Package bybit
package bybit

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/govalues/decimal"
	"github.com/hirokisan/bybit/v2"
	"github.com/strike-finance/strike-v2-backend/services/price/config"
	"github.com/strike-finance/strike-v2-backend/services/price/exchange"
)

type Bybit struct {
	client             *bybit.Client
	config             *config.Config
	pairs              []string
	doneCh             chan struct{}
	currentEndpointIdx int
	pairsCache         map[string]bool // Optimize filter speed when fetch prices back from OKX
}

func New(config *config.Config) *Bybit {
	formattedPairs := formatPairs(config.Pairs)
	pairsCache := make(map[string]bool)
	for _, pair := range formattedPairs {
		pairsCache[pair] = true
	}
	return &Bybit{
		client:             bybit.NewClient(),
		config:             config,
		pairs:              formattedPairs,
		doneCh:             make(chan struct{}),
		currentEndpointIdx: 0,
		pairsCache:         pairsCache,
	}
}

func (b *Bybit) GetName() string {
	return "Bybit"
}

func (b *Bybit) Active(priceFeedCh chan<- exchange.PriceFeed) {
	ticker := time.NewTicker(exchange.PriceFeedInterval)
	for {
		select {
		case <-b.doneCh:
			log.Info("Bybit Active stopped")
			return
		default:
		}

		endpoint := b.getEndpoint()
		b.client.WithBaseURL(endpoint)

		resp, err := b.client.V5().Market().GetTickers(bybit.V5GetTickersParam{
			Category: bybit.CategoryV5Spot,
		})
		if err != nil {
			log.Error(err)
			continue
		}
		for _, item := range resp.Result.Spot.List {
			// Filter out pairs not in config
			if !b.pairsCache[string(item.Symbol)] {
				continue
			}

			price, _ := decimal.Parse(item.LastPrice)
			priceFeedCh <- exchange.PriceFeed{
				Exchange:    b.GetName(),
				Pair:        string(item.Symbol),
				TimeInMilli: time.Now().UnixMilli(),
				Price:       price,
			}
		}

		<-ticker.C
	}
}

func (b *Bybit) Shutdown() {
	select {
	case <-b.doneCh:
	default:
		close(b.doneCh)
	}
}
