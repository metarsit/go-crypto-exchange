package method

import (
	"net/url"
	"testing"

	"gotest.tools/assert"
)

func TestPostHTTPGoogleError(t *testing.T) {
	url, _ := url.Parse("www.google.com")
	_, err := Post(url, nil)
	assert.Error(t, err, `Post "www.google.com": unsupported protocol scheme ""`)
}

func TestPostBadGateway(t *testing.T) {
	url, _ := url.Parse("https://api.crypto.com/v1/account")
	result, err := Post(url, nil)
	assert.NilError(t, err)
	assert.Equal(t, result.StatusCode, 502)
	assert.Equal(t, result.Status, "502 Bad Gateway")
}
