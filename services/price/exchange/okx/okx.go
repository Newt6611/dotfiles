// Package mexc
package okx

import (
	"strconv"
	"time"

	"github.com/charmbracelet/log"
	"github.com/govalues/decimal"
	"github.com/strike-finance/strike-v2-backend/services/price/config"
	"github.com/strike-finance/strike-v2-backend/services/price/exchange"
)

type OKX struct {
	config     *config.Config
	doneCh     chan struct{}
	pairsCache map[string]bool // Optimize filter speed when fetch prices back from OKX
}

func New(config *config.Config) *OKX {
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

func (b *OKX) Active(priceFeedCh chan<- exchange.PriceFeed) {
	ticker := time.NewTicker(exchange.PriceFeedInterval)
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
			if !b.pairsCache[ticker.InstID] {
				continue
			}

			price, _ := decimal.Parse(ticker.Last)
			ts, _ := strconv.ParseInt(ticker.TS, 10, 64)
			priceFeedCh <- exchange.PriceFeed{
				Exchange:    b.GetName(),
				Pair:        ticker.InstID,
				TimeInMilli: ts,
				Price:       price,
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
