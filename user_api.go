package exchange

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strings"

	"github.com/metarsit/exchange/internal/method"
)

// Account lists all available market symbols
func (u *UserAuth) Account() (UserResponse, error) {
	accountURL := URL("/v1/account")
	values := url.Values{
		"api_key": []string{u.APIKey},
		"time":    []string{u.Time},
		"sign":    []string{u.Sign},
	}
	var result UserResponse
	resp, err := method.Post(accountURL, strings.NewReader(values.Encode()))
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
