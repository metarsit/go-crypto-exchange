package exchange

import (
	"encoding/json"
	"io/ioutil"

	"github.com/metarsit/exchange/internal/method"
)

// Symbols return the whole list of market symbols
func (m *Market) Symbols() (MarketResponse, error) {
	SymbolURL := URL("/v1/symbols")

	var result MarketResponse
	resp, err := method.Get(SymbolURL, nil)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return result, err
	}
	return result, nil
}
