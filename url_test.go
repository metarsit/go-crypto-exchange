package exchange

import (
	"net/url"
	"testing"

	"gotest.tools/assert"
)

func TestURL(t *testing.T) {
	result := URL("path")
	expect := &url.URL{
		Scheme:     "https",
		Host:       "api.crypto.com",
		ForceQuery: false,
		Path:       "path",
	}

	assert.DeepEqual(t, result, expect)
}
