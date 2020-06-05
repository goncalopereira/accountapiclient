package e2e_test

import (
	"github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPost_WhenAPIConnectErrorThenReturnError(t *testing.T) {
	r := http.NewRequest()
	response, err := r.Post("http://127.0.0.1:111", nil) //use IPV4 here otherwise you might get IPV6

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestGet_WhenAPIConnectErrorThenReturnError(t *testing.T) {
	r := http.NewRequest()
	response, err := r.Get("http://127.0.0.1:111") //use IPV4 here otherwise you might get IPV6

	assert.NotNil(t, err)
	assert.Nil(t, response)
}
