//nolint:lll
package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/goncalopereira/accountapiclient/internal/data"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/internal/test"
	"github.com/stretchr/testify/assert"
)

//createResponseData returns valid static Data with an Account.
func createResponseData() *data.Data {
	account := test.NewAccountDataFromFile("create-response.json")
	return account
}

func NewServerWithResponse(response *internalhttp.Response) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(response.StatusCode)
		fmt.Fprintln(w, string(response.Body))
	}))

	return ts
}

func TestGet_WhenResponseIsOKThenStatusOKAndReturnBody(t *testing.T) {
	originalResponse := &internalhttp.Response{StatusCode: http.StatusOK, Body: test.ReadJSON("fetch-response.json")}
	ts := NewServerWithResponse(originalResponse)

	r := internalhttp.NewClient()

	req, err := http.NewRequest("GET", ts.URL, nil)
	assert.Nil(t, err)

	response, err := r.Do(req)

	assert.Nil(t, err)
	assert.Equal(t, originalResponse.StatusCode, response.StatusCode)
	//http server response ends with \n - harder to compare responses
	assert.Equal(t, strings.ReplaceAll(string(originalResponse.Body), "\n", ""), strings.ReplaceAll(string(response.Body), "\n", ""))
}

func TestPost_WhenDataSentAndResponseIsOKThenStatusOKAndReturnBody(t *testing.T) {
	originalResponse := &internalhttp.Response{StatusCode: http.StatusCreated, Body: test.ReadJSON("create-response.json")}
	ts := NewServerWithResponse(originalResponse)

	req, err := http.NewRequest("POST", ts.URL, bytes.NewBuffer(test.ReadJSON("create-request.json")))
	assert.Nil(t, err)

	response, err := internalhttp.NewClient().Do(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, response.StatusCode)

	var result = data.Data{}
	err = json.Unmarshal(response.Body, &result)

	assert.Nil(t, err)
	assert.Equal(t, createResponseData().Account, result.Account)
}

func TestDelete_WhenDataSentAndResponseIsOKThenStatusOKAndReturnBody(t *testing.T) {
	ts := NewServerWithResponse(&internalhttp.Response{StatusCode: http.StatusNoContent, Body: nil})

	req, err := http.NewRequest("DELETE", ts.URL, nil)
	assert.Nil(t, err)

	response, err := internalhttp.NewClient().Do(req)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}
