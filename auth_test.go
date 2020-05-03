package exchange

import (
	"testing"

	"gotest.tools/assert"
)

func TestNewUserAPI(t *testing.T) {
	var (
		api    = "xxx"
		secret = "yyy"
	)
	result, err := NewUserAPI(api, secret)
	assert.NilError(t, err)

	expect := &UserAuth{
		APIKey:    "xxx",
		SecretKey: "yyy",
	}
	assert.DeepEqual(t, result, expect)
}

func TestEmptyKeyNewUserAPI(t *testing.T) {
	_, err := NewUserAPI("", "")
	assert.Error(t, err, "Keys cannot be empty")
}
