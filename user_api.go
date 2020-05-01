package exchange

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strings"
	"time"

	"github.com/metarsit/exchange/internal/method"
)

// Account lists all available market symbols
func (u *UserAuth) Account() (UserResponse, error) {
	accountURL := URL("/v1/account")

	var result UserResponse
	values, err := reqValues(u.APIKey, u.SecretKey, nil)
	if err != nil {
		return result, err
	}

	resp, err := method.Post(accountURL, strings.NewReader(values.Encode()))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	return bodyToUserResponse(resp.Body, &result)
}

// ShowOrder returns all available orders
func (u *UserAuth) ShowOrder(id, symbol string) (UserResponse, error) {
	showOrderURL := URL("/v1/showOrder")

	values, err := reqValues(u.APIKey, u.SecretKey, map[string]string{
		"order_id": id,
		"symbol":   symbol,
	})

	var result UserResponse
	resp, err := method.Post(showOrderURL, strings.NewReader(values.Encode()))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	return bodyToUserResponse(resp.Body, &result)
}

func bodyToUserResponse(body io.Reader, result *UserResponse) (UserResponse, error) {
	respBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return *result, err
	}
	if err := json.Unmarshal(respBytes, &result); err != nil {
		return *result, err
	}
	return *result, nil
}

func reqValues(apiKey, secretKey string, param map[string]string) (url.Values, error) {
	values := url.Values{}
	// Time
	time := fmt.Sprint(time.Now().UnixNano() / 1000000)

	// Parameters
	paramStr := ""
	if param != nil {
		for key, value := range param {
			paramStr += (key + value)
			values.Add(key, value)
		}
	}

	// Sign
	preEncodeMsg := fmt.Sprintf(`api_key%s%stime%s%s`,
		apiKey,
		paramStr,
		time,
		secretKey,
	)
	encode := sha256.New()
	_, err := encode.Write([]byte(preEncodeMsg))
	if err != nil {
		return nil, err
	}
	sign := hex.EncodeToString(encode.Sum(nil))

	values.Set("api_key", apiKey)
	values.Set("time", time)
	values.Set("sign", sign)

	return values, nil
}
