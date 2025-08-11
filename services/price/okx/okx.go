// Package okx
package okx

import (
	"strconv"
	"time"

	"github.com/charmbracelet/log"
	"github.com/govalues/decimal"
	"github.com/strike-finance/strike-v2-backend/services/price"
)

type OKX struct {
	config     *price.Config
	doneCh     chan struct{}
	pairsCache map[string]bool // Optimize filter speed when fetch prices back from OKX
}

func New(config *price.Config) *OKX {
	pairsCache := make(map[string]bool)
	for _, pair := range config.Pairs {
		pairsCache[pair] = true
	}

	return &OKX{
		config:     config,
		doneCh:     make(chan struct{}),
		pairsCache: pairsCache,
	}
}

func (b *OKX) GetName() string {
	return "OKX"
}

func (b *OKX) Active(priceFeedCh chan<- price.PriceFeed) {
	ticker := time.NewTicker(price.PriceFeedInterval)
	for {
		select {
		case <-b.doneCh:
			log.Info("OKX Active stopped")
			return
		default:
		}

		resp, err := b.MarketTickers()
		if err != nil {
			log.Error(err)
			continue
		}
		if resp.Code != "0" {
			log.Error(resp.Msg)
			continue
		}

		for _, ticker := range resp.Data {
			// Filter out pairs not in config
			if !b.pairsCache[ticker.InstID] {
				continue
			}

			newPrice, _ := decimal.Parse(ticker.Last)
			ts, _ := strconv.ParseInt(ticker.TS, 10, 64)
			priceFeedCh <- price.PriceFeed{
				Exchange:    b.GetName(),
				Pair:        ticker.InstID,
				TimeInMilli: ts,
				Price:       newPrice,
			}
		}

		<-ticker.C
	}
}

func (b *OKX) Shutdown() {
	select {
	case <-b.doneCh:
	default:
		close(b.doneCh)
	}
}

func (b *OKX) GetEndpoint() string {
	return b.config.OKXEndpoint
}
