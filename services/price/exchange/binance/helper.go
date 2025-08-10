package binance

import "strings"

func formatPair(pairs []string) []string {
	correctPairs := make([]string, len(pairs))

	// Pair from config is BTCUSDT
	// Parse it to btcusdt
	for i, pair := range pairs {
		newPair := strings.ToLower(pair)
		correctPairs[i] = newPair + "@" + string(EventTypeAggTrade)
	}

	return correctPairs
}
