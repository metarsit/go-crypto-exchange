package exchange

import (
	"encoding/json"
)

// UserAuth contains required necessary information
// to hit REST APIs of Crypto Exchange
type UserAuth struct {
	APIKey string `json:"api_key"`
	Time   string `json:"time"`
	Sign   string `json:"sign"`
}

// UserResponse is Crypto Exchange Response struct
type UserResponse struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
	Data    Data   `json:"data"`
}

// Data contains total asset and list of coins' information
type Data struct {
	TotalAsset string `json:"total_asset"`
	CoinList   []Coin `json:"coin_list"`
}

// Coin contains coin's information
type Coin struct {
	Normal       string `json:"normal"`
	Locked       string `json:"locked"`
	BTCValuation string `json:"btcValuation"`
	Name         string `json:"coin"`
}

// Market is a struct for Crypto Exchange
// Market API
type Market struct{}

// MarketResponse is Crypto Exchange Response of market struct
type MarketResponse struct {
	Code    string           `json:"code"`
	Message string           `json:"msg"`
	Data    *json.RawMessage `json:"data"`
}

// MarketSymbols is information of the symbol
type MarketSymbols struct {
	Name            string `json:"symbol"`
	CountCoin       string `json:"count_coin"`
	AmountPrecision int    `json:"amount_precision"`
	BaseCoin        string `json:"base_coin"`
	PricePrecision  int    `json:"price_precision"`
}

// MarketDepth contains a Tick of Symbol's depth
type MarketDepth struct {
	Tick Tick `json:"tick"`
}

// Tick consists of slice of asks, bids and server time
type Tick struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
	Time int64       `json:"time"`
}
