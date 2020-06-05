package e2e

import (
	"github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestPost_WhenAPIConnectErrorThenReturnError(t *testing.T) {
	r := http.NewRequest()
	u, _ := url.Parse("http://127.0.0.1:111")
	response, err := r.Post(u, nil) //use IPV4 here otherwise you might get IPV6

	assert.NotNil(t, err)
	assert.Nil(t, response)
}

func TestGet_WhenAPIConnectErrorThenReturnError(t *testing.T) {
	r := http.NewRequest()
	u, _ := url.Parse("http://127.0.0.1:111")
	response, err := r.Get(u) //use IPV4 here otherwise you might get IPV6

	assert.NotNil(t, err)
	assert.Nil(t, response)
}
