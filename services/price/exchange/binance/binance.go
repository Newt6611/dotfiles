// Package binance
package binance

import (
	"encoding/json"
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/gorilla/websocket"
	"github.com/strike-finance/strike-v2-backend/services/price/config"
	"github.com/strike-finance/strike-v2-backend/services/price/exchange"
)

type Binance struct {
	config             *config.Config
	pairs              []string
	doneCh             chan struct{}
	currentEndpointIdx int
}

func New(config *config.Config) *Binance {
	return &Binance{
		config:             config,
		pairs:              formatPair(config.Pairs),
		doneCh:             make(chan struct{}),
		currentEndpointIdx: 0,
	}
}

func (b *Binance) GetName() string {
	return "Binance"
}

func (b *Binance) Active(priceFeedCh chan<- exchange.PriceFeed) {
	for {
		select {
		case <-b.doneCh:
			log.Info("Binance Active stopped")
			return
		default:
		}

		endpoint := b.getWssEndpoint()
		c, _, err := websocket.DefaultDialer.Dial(endpoint, nil)
		if err != nil {
			exchange.WsReconnectWait(b.GetName(), "Dial: "+err.Error())
			continue
		}

		// Send subscription JSON
		subMsg := Subscribe{
			Method: "SUBSCRIBE",
			Params: b.pairs,
			ID:     1,
		}

		if err := c.WriteJSON(subMsg); err != nil {
			c.Close()
			exchange.WsReconnectWait(b.GetName(), "Subscribe: "+err.Error())
			continue
		}

		for {
			select {
			case <-b.doneCh:
				log.Info("Binance price feed loop stopped")
				c.Close()
				return
			default:
			}

			_, respMsg, err := c.ReadMessage()
			if err != nil {
				c.Close()
				exchange.WsReconnectWait(b.GetName(), "Read: "+err.Error())
				goto reconnect
			}

			var aggTrade AggTrade
			err = json.Unmarshal(respMsg, &aggTrade)
			if err != nil {
				log.Warn("Binance unmarshal:", err)
				continue
			}

			if aggTrade.EventType != string(EventTypeAggTrade) {
				continue
			}

			price, _ := strconv.ParseFloat(aggTrade.Price, 64)
			priceFeedCh <- exchange.PriceFeed{
				Exchange:    b.GetName(),
				Pair:        aggTrade.Symbol,
				TimeInMilli: aggTrade.EventTime,
				Price:       price,
			}
		}
	reconnect:
	}
}

func (b *Binance) Shutdown() {
	select {
	case <-b.doneCh:
	default:
		close(b.doneCh)
	}
}

func (b *Binance) getWssEndpoint() string {
	b.currentEndpointIdx++
	if b.currentEndpointIdx >= len(b.config.BinanceWss) {
		b.currentEndpointIdx = 0
	}
	return b.config.BinanceWss[b.currentEndpointIdx]
}
