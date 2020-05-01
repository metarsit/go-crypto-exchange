package exchange

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// NewUserAPI creates an instance for User
func NewUserAPI(apiKey, secretKey string, time int64) (*UserAuth, error) {
	preEncodeMsg := fmt.Sprintf(`api_key%stime%s%s`,
		apiKey,
		fmt.Sprint(time),
		secretKey,
	)
	encode := sha256.New()
	_, err := encode.Write([]byte(preEncodeMsg))
	if err != nil {
		return nil, err
	}
	return &UserAuth{
		APIKey: apiKey,
		Time:   fmt.Sprint(time),
		Sign:   hex.EncodeToString(encode.Sum(nil)),
	}, nil
}

// NewMarketAPI creates an instance for Market
func NewMarketAPI() *Market {
	return &Market{}
}
