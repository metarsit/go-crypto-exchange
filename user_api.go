package exchange

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/metarsit/exchange/internal/method"
)

// Account lists all available market symbols
func (u *UserAuth) Account() (UserResponse, error) {
	accountURL := URL("/v1/account")
	var result UserResponse

	values, err := reqValues(u.SecretKey, map[string]string{
		"api_key": u.APIKey,
	})
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
	var result UserResponse

	values, err := reqValues(u.SecretKey, map[string]string{
		"api_key":  u.APIKey,
		"order_id": id,
		"symbol":   symbol,
	})
	if err != nil {
		return result, err
	}

	resp, err := method.Post(showOrderURL, strings.NewReader(values.Encode()))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	return bodyToUserResponse(resp.Body, &result)
}

// AllOrders lists all orders in a particular market
func (u *UserAuth) AllOrders(symbol, startDate, endDate string, page, pageSize int) (UserResponse, error) {
	allOrders := URL("/v1/allOrders")
	var result UserResponse

	pattern := regexp.MustCompile(`[0-9]{4}-[0-1][0-9]-[0-3][0-9] [0-2][0-9]:[0-5][0-9]:[0-5][0-9]`)
	if !pattern.Match([]byte(startDate)) || !pattern.Match([]byte(endDate)) {
		return result, errors.New(`the input date does not match the format: YYYY:MM:DD HH:mm:ss`)
	}

	values, err := reqValues(u.SecretKey, map[string]string{
		"api_key":   u.APIKey,
		"endDate":   endDate,
		"page":      fmt.Sprint(page),
		"pageSize":  fmt.Sprint(pageSize),
		"startDate": startDate,
		"symbol":    symbol,
	})
	if err != nil {
		return result, err
	}

	resp, err := method.Post(allOrders, strings.NewReader(values.Encode()))
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	return bodyToUserResponse(resp.Body, &result)
}

// CancelOrder cancels a specified order
func (u *UserAuth) CancelOrder(symbol string, orderID int) (UserResponse, error) {
	allOrders := URL("/v1/orders/cancel")
	var result UserResponse

	values, err := reqValues(u.SecretKey, map[string]string{
		"api_key":  u.APIKey,
		"order_id": fmt.Sprint(orderID),
		"symbol":   symbol,
	})
	if err != nil {
		return result, err
	}

	resp, err := method.Post(allOrders, strings.NewReader(values.Encode()))
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

func reqValues(secretKey string, param map[string]string) (url.Values, error) {
	values := url.Values{}

	if param == nil {
		return nil, errors.New("parameter cannot be nil")
	}

	// Time
	param["time"] = fmt.Sprint(time.Now().UnixNano() / 1000000)

	// Sorted Parameters
	keys := []string{}
	paramStr := ""
	for key := range param {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, value := range keys {
		paramStr += (value + param[value])
		values.Add(value, param[value])
	}

	// Sign
	preEncodeMsg := paramStr + secretKey
	encode := sha256.New()
	_, err := encode.Write([]byte(preEncodeMsg))
	if err != nil {
		return nil, err
	}
	sign := hex.EncodeToString(encode.Sum(nil))
	values.Set("sign", sign)

	return values, nil
}
