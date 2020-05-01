package method

import (
	"net/url"
	"testing"

	"gotest.tools/assert"
)

func TestGetHTTPGoogleError(t *testing.T) {
	url, _ := url.Parse("www.google.com")
	_, err := Get(url, nil)
	assert.Error(t, err, `Get "www.google.com": unsupported protocol scheme ""`)
}

func TestGet(t *testing.T) {
	url, _ := url.Parse("https://google.com")
	result, err := Get(url, nil)
	assert.NilError(t, err)
	assert.Equal(t, result.StatusCode, 200)
	assert.Equal(t, result.Status, "200 OK")
}
