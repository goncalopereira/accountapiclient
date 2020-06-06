package e2e_test

import (
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestPost_WhenAPIConnectErrorThenReturnError(t *testing.T) {
	r := internalhttp.NewRequest()

	req, _ := http.NewRequest("POST", "http://127.0.0.1:111", nil)
	response, err := r.Do(req) //use IPV4 here otherwise you might get IPV6

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestGet_WhenAPIConnectErrorThenReturnError(t *testing.T) {
	r := internalhttp.NewRequest()

	req, _ := http.NewRequest("GET", "http://127.0.0.1:111", nil)
	response, err := r.Do(req) //use IPV4 here otherwise you might get IPV6

	assert.NotNil(t, err)
	assert.Nil(t, response)
}
