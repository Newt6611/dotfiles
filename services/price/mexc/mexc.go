// Package mexc
package mexc

import (
	"time"

	"github.com/charmbracelet/log"
	"github.com/govalues/decimal"
	"github.com/strike-finance/strike-v2-backend/services/price"
)

type Mexc struct {
	config     *price.Config
	pairs      []string
	doneCh     chan struct{}
	pairsCache map[string]bool // Optimize filter speed when fetch prices back from OKX
}

func New(config *price.Config) *Mexc {
	formattedPairs := formatPairs(config.Pairs)
	pairsCache := make(map[string]bool)
	for _, pair := range formattedPairs {
		pairsCache[pair] = true
	}
	return &Mexc{
		config:     config,
		pairs:      formattedPairs,
		pairsCache: pairsCache,
		doneCh:     make(chan struct{}),
	}
}

func (b *Mexc) GetName() string {
	return "Mexc"
}

func (b *Mexc) Active(priceFeedCh chan<- price.PriceFeed) {
	ticker := time.NewTicker(price.PriceFeedInterval)
	for {
		select {
		case <-b.doneCh:
			log.Info("Mexc Active stopped")
			return
		default:
		}

		symbols, err := b.TickerPricese()
		if err != nil {
			log.Error(err)
			continue
		}

		for _, symbol := range symbols {
			// Filter out pairs not in config
			if !b.pairsCache[symbol.Symbol] {
				continue
			}

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

func (b *Mexc) Shutdown() {
	select {
	case <-b.doneCh:
	default:
		close(b.doneCh)
	}
}
