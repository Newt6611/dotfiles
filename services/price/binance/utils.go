package binance

import "strings"

func formatPairs(pairs []string) []string {
	var formattedPairs []string
	for _, pair := range pairs {
		newPair := formatPair(pair)
		formattedPairs = append(formattedPairs, newPair)
	}
	return formattedPairs
}

func formatPair(pair string) string {
	return strings.ToUpper(strings.ReplaceAll(pair, "-", ""))
}

func (b *Binance) getEndpoint() string {
	b.currentEndpointIdx++
	if b.currentEndpointIdx >= len(b.config.BinanceEndpoints) {
		b.currentEndpointIdx = 0
	}
	return b.config.BinanceEndpoints[b.currentEndpointIdx]
}
