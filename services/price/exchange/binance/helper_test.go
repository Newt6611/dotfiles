package binance

import "testing"

func TestFormatPair(t *testing.T) {
	pairs := []string { "BTCUSDT", "ETHUSDT", "ADAUSDT" }
	expectedPairs := []string { "btcusdt@aggTrade", "ethusdt@aggTrade", "adausdt@aggTrade" }
	result := formatPair(pairs)
	for i, pair := range result {
		if pair != expectedPairs[i] {
			t.Errorf("Expected %s, got %s", expectedPairs[i], pair)
		}
	}
}
