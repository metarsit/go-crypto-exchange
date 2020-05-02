package exchange

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
	"time"

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

// Ticker gets ticker for a/or particular market
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

// Trades gets the last 200 trades in a specified market
func (m *Market) Trades(symbol string) (MarketResponse, error) {
	tradesURL := URL("/v1/trades")
	var result MarketResponse

	if symbol == "" {
		return result, errors.New("Symbol cannot be empty")
	}

	query := url.Values{
		"symbol": []string{symbol},
	}

	resp, err := method.Get(tradesURL, nil, query)
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

// KLines gets k-line data over a specified period
func (m *Market) KLines(symbol string, period int) (MarketResponse, error) {
	kLinesURL := URL("/v1/klines")
	var result MarketResponse

	if symbol == "" {
		return result, errors.New("Symbol cannot be empty")
	}

	if !allowedPeriod(period) {
		return result, errors.New("Period selected is not allowed")
	}

	query := url.Values{
		"period": []string{strconv.Itoa(period)},
		"symbol": []string{symbol},
	}

	resp, err := method.Get(kLinesURL, nil, query)
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

func allowedPeriod(period int) bool {
	for _, allowed := range []time.Duration{1, 5, 15, 30, 60, 1440, 10080, 43200} {
		duration, err := time.ParseDuration(fmt.Sprintf("%dm", period))
		if err != nil {
			continue
		}
		if (allowed * time.Minute) == duration {
			return true
		}
	}
	return false
}
