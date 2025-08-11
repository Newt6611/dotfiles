package mexc

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
