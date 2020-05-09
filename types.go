package exchange

import (
	"encoding/json"
)

// UserAuth contains required necessary information
// to hit REST APIs of Crypto Exchange
type UserAuth struct {
	APIKey    string
	SecretKey string
}

// UserResponse is Crypto Exchange Response struct
type UserResponse struct {
	Code    string           `json:"code"`
	Message string           `json:"msg"`
	Data    *json.RawMessage `json:"data"`
}

// Balance contains total asset and list of coins' information
type Balance struct {
	TotalAsset string `json:"total_asset"`
	CoinList   []Coin `json:"coin_list"`
}

// Orders contains trade information of order(s)
type Orders struct {
	OrderInfo OrderInfo `json:"order_info"`
}

// AllOrders contains all requested order information
type AllOrders struct {
	Count     int         `json:"count"`
	OrderList []OrderInfo `json:"orderList"`
}

// MyTrades contains all trades information
type MyTrades struct {
	Count      int         `json:"count"`
	ResultList []UserTrade `json:"resultList"`
}

// Coin contains coin's information
type Coin struct {
	Normal       string `json:"normal"`
	Locked       string `json:"locked"`
	BTCValuation string `json:"btcValuation"`
	Name         string `json:"coin"`
}

// OrderInfo contains an order information
type OrderInfo struct {
	AvgPrice     string      `json:"avg_price"`
	BaseCoin     string      `json:"baseCoin"`
	CountCoin    string      `json:"countCoin"`
	CreatedAt    int64       `json:"created_at"`
	DealPrice    float64     `json:"deal_price"`
	DealVolume   string      `json:"deal_volume"`
	Fee          string      `json:"fee"`
	FeeCoin      string      `json:"fee_coin"`
	ID           int         `json:"id"`
	Price        string      `json:"price"`
	RemainVolume string      `json:"remain_volume"`
	Side         string      `json:"side"`
	SideMsg      string      `json:"side_msg"`
	Source       int         `json:"source"`
	SourceMsg    string      `json:"source_msg"`
	Status       int         `json:"status"`
	StatusMsg    string      `json:"status_msg"`
	TotalPrice   string      `json:"total_price"`
	TradeList    []UserTrade `json:"tradeList"`
	Type         int         `json:"type"`
	Volume       string      `json:"volume"`
}

// UserTrade contains an information of the trade
type UserTrade struct {
	AskID     int64  `json:"ask_id,omitempty"`
	AskUserID int64  `json:"ask_user_id,omitempty"`
	BidID     int64  `json:"bid_id,omitempty"`
	BidUserID int64  `json:"bid_user_id,omitempty"`
	CTime     int64  `json:"ctime"`
	DealPrice string `json:"deal_price"`
	Fee       string `json:"fee"`
	FeeCoin   string `json:"feeCoin"`
	ID        int    `json:"id"`
	Price     string `json:"price"`
	Type      string `json:"type"`
	Volume    string `json:"volume"`
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

// TickerPrice contains all prices
type TickerPrice map[string]float64

// TickerAll contains server time and list of all tickers
type TickerAll struct {
	Date int64    `json:"date"`
	List []Symbol `json:"ticker"`
}

// Symbol contains ticker's information
type Symbol struct {
	Name   string `json:"symbol"`
	High   string `json:"high"`
	Low    string `json:"low"`
	Volume string `json:"vol"`
	Last   string `json:"last"`
	Buy    string `json:"buy"`
	Sell   string `json:"sell"`
	Change string `json:"change"`
	Rose   string `json:"rose"`
}

// MarketTrade contains the information of the trade
// of specific market at the point of time
type MarketTrade struct {
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
	CTime  int64   `json:"ctime"`
	ID     int     `json:"id"`
	Type   string  `json:"type"`
}
