// Package models
package models

import "github.com/govalues/decimal"

type PriceEvent struct {
	Exchange  string          `json:"exchange"`
	Pair      string          `json:"pair"`
	Price     decimal.Decimal `json:"price"`
	Timestamp int64           `json:"timestamp"`
}
