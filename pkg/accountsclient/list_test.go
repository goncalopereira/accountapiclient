//nolint:scopelint,funlen
package accountsclient_test

import (
	"github.com/goncalopereira/accountapiclient/internal/config"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"github.com/goncalopereira/accountapiclient/pkg/accountsclient"
	"github.com/goncalopereira/accountapiclient/test"
	configtest "github.com/goncalopereira/accountapiclient/test/config"
	httptest "github.com/goncalopereira/accountapiclient/test/http"
	"github.com/stretchr/testify/assert"
	"net/url"
	"reflect"
	"testing"
)

func TestClient_List(t *testing.T) {
	type fields struct {
		config  config.IAPI
		request http.IRequest
	}

	type args struct {
		urls *url.Values
	}

	multipleAccounts := test.NewAccountsFromFile("list-response.json")

	accountsBody, err := json.DataToBytes(multipleAccounts)
	assert.Nil(t, err)

	accountsResponse := &http.Response{StatusCode: 200, Body: accountsBody}

	emptyList := test.NewAccountsFromFile("list-empty.json")

	emptyAccountsBody, err := json.DataToBytes(emptyList)
	assert.Nil(t, err)

	emptyAccountsResponse := &http.Response{StatusCode: 200, Body: emptyAccountsBody}

	apiErrorMessage := test.NewErrorMessageFromFile("server-error.json")

	errorBody, err := json.DataToBytes(apiErrorMessage)
	assert.Nil(t, err)

	errorResponse := &http.Response{StatusCode: 500, Body: errorBody}

	brokenResponse := &http.Response{StatusCode: 500, Body: nil}

	api := config.DefaultAPI()
	brokenAPI := configtest.NewAPIMock(nil, test.ErrBrokenConfig)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    data.IOutput
		wantErr bool
	}{
		{"GivenAccountsWhenDefaultQueryThenReturnAccounts",
			fields{config: api, request: httptest.NewGetRequestMock(accountsResponse, nil)},
			args{urls: &url.Values{}},
			multipleAccounts,
			false},
		{"GivenNoAccountsWhenDefaultQueryThenReturnNilArray",
			fields{config: api, request: httptest.NewGetRequestMock(emptyAccountsResponse, nil)},
			args{urls: &url.Values{}},
			&account.AccountsData{},
			false},
		{"WhenNon200ThenReturnErrorMessage",
			fields{config: api, request: httptest.NewGetRequestMock(errorResponse, nil)},
			args{urls: &url.Values{}},
			apiErrorMessage,
			false},
		{"WhenNon200BrokenResponseThenReturnError",
			fields{config: api, request: httptest.NewGetRequestMock(brokenResponse, nil)},
			args{urls: &url.Values{}},
			&data.NoOp{},
			true},
		{"WhenHTTPClientThrowsThenReturnError",
			fields{config: api, request: httptest.NewGetRequestMock(nil, test.ErrBrokenHTTPClient)},
			args{urls: &url.Values{}},
			&data.NoOp{},
			true},
		{"WhenBrokenAPIConfigThrowsThenReturnError",
			fields{config: brokenAPI, request: nil},
			args{urls: &url.Values{}},
			&data.NoOp{},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := accountsclient.NewClient(tt.fields.config, tt.fields.request)
			got, err := c.List(tt.args.urls)
			if (err != nil) != tt.wantErr {
				t.Errorf("List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) && tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List() got = %v, want %v", got, tt.want)
			}
		})
	}
}
