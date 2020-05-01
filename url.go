package exchange

import "net/url"

// URL is URL struct of Crypto Exchange API
func URL(path string) *url.URL {
	return &url.URL{
		Scheme:     "https",
		Host:       "api.crypto.com",
		ForceQuery: false,
		Path:       path,
	}
}
