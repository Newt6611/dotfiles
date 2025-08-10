// Package models
package models

type PriceEvent struct {
	Exchange  string  `json:"exchange"`
	Pair      string  `json:"pair"`
	Price     float64 `json:"price"`
	Timestamp string  `json:"timestamp"`
}
