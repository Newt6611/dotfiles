// Package exchange provides an exchange interface.
package exchange

import "time"

type ExchangeType string

const (
	Binance ExchangeType = "Binance"
)

const (
	ReconnectInterval = time.Second * 5
)

type PriceFeed struct {
	Exchange    string
	Pair        string
	TimeInMilli int64
	Price       float64
}

type Exchange interface {
	GetName() string
	Active(priceFeedCh chan<- PriceFeed)
	Shutdown()
}
