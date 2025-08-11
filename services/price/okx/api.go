package okx

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (b *OKX) MarketTickers() (*MarketTickerResp, error) {
	endpint := fmt.Sprintf("%s%s?instType=SPOT", b.config.OKXEndpoint, "/api/v5/market/tickers")
	resp, err := http.Get(endpint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result MarketTickerResp
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
