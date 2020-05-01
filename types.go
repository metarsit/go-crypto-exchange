package exchange

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
