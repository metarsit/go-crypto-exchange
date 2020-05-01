package method

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

const timeout = time.Second * 5

// Post is a customized POST method to hit crypto.com exchange API
func Post(url *url.URL, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest("POST", url.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: timeout}
	return client.Do(req)
}
