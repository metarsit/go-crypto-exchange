package exchange

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"

	"github.com/metarsit/exchange/internal/method"
)

// Symbols lists all available market symbols
func (m *Market) Symbols() (MarketResponse, error) {
	symbolURL := URL("/v1/symbols")

	var result MarketResponse
	resp, err := method.Get(symbolURL, nil, nil)
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

// Depth gets the order book for a particular market
func (m *Market) Depth(symbol, option string) (MarketResponse, error) {
	depthURL := URL("/v1/depth")
	var result MarketResponse

	if symbol == "" || option == "" {
		return result, errors.New("Neither symbol nor option can be empty")
	}

	query := url.Values{
		"symbol": []string{symbol},
		"type":   []string{option},
	}

	resp, err := method.Get(depthURL, nil, query)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// Crypto.com Exchange does not return 404 when symbol does not show
	if resp.StatusCode == 500 {
		return result, errors.New("Symbol does not exist")
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return result, err
	}
	return result, nil
}

// TickerPrice returns latest execution price for all markets
func (m *Market) TickerPrice() (MarketResponse, error) {
	tickerPriceURL := URL("/v1/ticker/price")
	var result MarketResponse

	resp, err := method.Get(tickerPriceURL, nil, nil)
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

// Ticker gets last 200 trades in a specified market
func (m *Market) Ticker(symbol string) (MarketResponse, error) {
	tickerURL := URL("/v1/ticker")
	var result MarketResponse
	values := url.Values{}

	if symbol != "" {
		values.Add("symbol", symbol)
	}

	resp, err := method.Get(tickerURL, nil, values)
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
