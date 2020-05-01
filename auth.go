package exchange

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// NewUserAPI creates a instance that
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
