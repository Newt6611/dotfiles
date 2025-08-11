package mexc

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (b *Mexc) TickerPricese() ([]SymbolPrice, error) {
	endpint := fmt.Sprintf("%s%s?symbols=all", b.config.MexcEndpoint, "/api/v3/ticker/price")
	resp, err := http.Get(endpint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []SymbolPrice
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
