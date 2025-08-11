package bybit

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

func (b *Bybit) getEndpoint() string {
	b.currentEndpointIdx++
	if b.currentEndpointIdx >= len(b.config.BybitEndpoints) {
		b.currentEndpointIdx = 0
	}
	return b.config.BybitEndpoints[b.currentEndpointIdx]
}
