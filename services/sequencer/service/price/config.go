// Package price
package price

import (
	"github.com/strike-finance/strike-v2-backend/services/price"
	"github.com/strike-finance/strike-v2-backend/services/sequencer/config"
)

func parseConfig(cfg *config.Config) price.Config {
	pairs := make([]string, len(cfg.Pairs))
	copy(pairs, cfg.Pairs)

	binanceEndpoints := make([]string, len(cfg.BinanceEndpoints))
	copy(binanceEndpoints, cfg.BinanceEndpoints)

	bybitEndpoints := make([]string, len(cfg.BybitEndpoints))
	copy(bybitEndpoints, cfg.BybitEndpoints)

	return price.Config{
		Pairs:            pairs,
		BinanceEndpoints: binanceEndpoints,
		OKXEndpoint:      cfg.OKXEndpoint,
		BybitEndpoints:   bybitEndpoints,
		MexcEndpoint:     cfg.MexcEndpoint,
	}
}
