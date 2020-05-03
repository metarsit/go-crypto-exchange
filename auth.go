package exchange

import "errors"

// NewUserAPI creates an instance for User
func NewUserAPI(apiKey, secretKey string) (*UserAuth, error) {
	if apiKey == "" || secretKey == "" {
		return nil, errors.New("Keys cannot be empty")
	}
	return &UserAuth{
		APIKey:    apiKey,
		SecretKey: secretKey,
	}, nil
}

// NewMarketAPI creates an instance for Market
func NewMarketAPI() *Market {
	return &Market{}
}
