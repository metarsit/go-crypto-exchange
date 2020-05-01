package method

import (
	"io"
	"net/http"
	"net/url"
)

// Get is a customized Get method to hit crypto.com exchange API
func Get(url *url.URL, body io.Reader, query url.Values) (*http.Response, error) {
	req, err := http.NewRequest("GET", url.String(), body)
	if err != nil {
		return nil, err
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: timeout}
	return client.Do(req)
}
