package exchange

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"strings"

	"github.com/metarsit/exchange/internal/method"
)

// Account lists all available market symbols
func (u *UserAuth) Account() (UserResponse, error) {
	accountURL := URL("/v1/account")

	var result UserResponse
	resp, err := method.Post(accountURL, strings.NewReader(u.Values.Encode()))
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
