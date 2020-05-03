package exchange

import (
	"fmt"
	"net/url"
	"testing"

	"gotest.tools/assert"
)

func TestNewUserAPI(t *testing.T) {
	var (
		time   int64 = 1588347414305
		api          = "xxx"
		secret       = "yyy"
	)
	result, err := NewUserAPI(api, secret, time)
	assert.NilError(t, err)

	expect := &UserAuth{
		Time:   fmt.Sprint(time),
		APIKey: api,
		Sign:   "a322d9b768b9d9cb08c5b205e69036660c61f1efe11ff9d35f222937c7994041",
		Values: &url.Values{
			"api_key": {"xxx"},
			"sign":    {"a322d9b768b9d9cb08c5b205e69036660c61f1efe11ff9d35f222937c7994041"},
			"time":    {"1588347414305"},
		},
	}
	assert.DeepEqual(t, result, expect)
}
