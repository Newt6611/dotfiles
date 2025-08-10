// Package exchange provides an exchange interface.
package exchange

import (
	"github.com/govalues/decimal"
	"time"
)

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
	Price       decimal.Decimal
}

type Exchange interface {
	GetName() string
	Active(priceFeedCh chan<- PriceFeed)
	Shutdown()
}
