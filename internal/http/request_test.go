//nolint:lll
package http_test

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"github.com/goncalopereira/accountapiclient/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func NewServerWithResponse(response *internalhttp.Response) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(response.StatusCode)
		fmt.Fprintln(w, string(response.Body))
	}))
	return ts
}

func TestGet_WhenResponseIsOKThenStatusOKAndReturnBody(t *testing.T) {
	originalResponse := internalhttp.NewResponse(http.StatusOK, test.ReadJSON("health-up.json"))
	ts := NewServerWithResponse(originalResponse)

	r := internalhttp.NewRequest()
	response, err := r.Get(ts.URL)

	assert.Nil(t, err)
	assert.Equal(t, originalResponse.StatusCode, response.StatusCode)
	//http server response ends with \n - harder to compare responses
	assert.Equal(t, strings.ReplaceAll(string(originalResponse.Body), "\n", ""), strings.ReplaceAll(string(response.Body), "\n", ""))
}

func TestPost_WhenDataSentAndResponseIsOKThenStatusOKAndReturnBody(t *testing.T) {
	ts := NewServerWithResponse(internalhttp.NewResponse(http.StatusCreated, test.ReadJSON("create-response.json")))

	requestBody, _ := json.DataToBody(test.ReadJSON("create.json"))
	response, err := internalhttp.NewRequest().Post(ts.URL, requestBody)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	var result = account.NewEmptyAccount()
	_ = json.BodyToData(response.Body, &result)
	assert.Equal(t, test.AccountCreateResponse().Data, result.Data)
}
