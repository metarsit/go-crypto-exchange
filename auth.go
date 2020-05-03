package exchange

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
)

// NewUserAPI creates an instance for User
func NewUserAPI(apiKey, secretKey string, time int64) (*UserAuth, error) {
	timeStr := fmt.Sprint(time)

	preEncodeMsg := fmt.Sprintf(`api_key%stime%s%s`,
		apiKey,
		timeStr,
		secretKey,
	)

	encode := sha256.New()
	_, err := encode.Write([]byte(preEncodeMsg))
	if err != nil {
		return nil, err
	}
	sign := hex.EncodeToString(encode.Sum(nil))

	return &UserAuth{
		APIKey: apiKey,
		Time:   timeStr,
		Sign:   sign,
		Values: &url.Values{
			"api_key": []string{apiKey},
			"time":    []string{timeStr},
			"sign":    []string{sign},
		},
	}, nil
}

// NewMarketAPI creates an instance for Market
func NewMarketAPI() *Market {
	return &Market{}
}
