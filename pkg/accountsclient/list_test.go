//nolint:scopelint,funlen
package accountsclient_test

import (
	"encoding/json"
	"github.com/goncalopereira/accountapiclient/internal/api"
	"github.com/goncalopereira/accountapiclient/internal/data"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/internal/test"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestClient_List(t *testing.T) {
	type fields struct {
		config  *api.API
		request internalhttp.IRequest
	}

	type args struct {
		urls *url.Values
	}

	multipleAccounts := test.NewAccountsFromFile("list-response.json")

	accountsBody, err := json.Marshal(multipleAccounts)
	assert.Nil(t, err)

	accountsResponse := &internalhttp.Response{StatusCode: http.StatusOK, Body: accountsBody}

	emptyList := test.NewAccountsFromFile("list-response-empty.json")

	emptyAccountsBody, err := json.Marshal(emptyList)
	assert.Nil(t, err)

	emptyAccountsResponse := &internalhttp.Response{StatusCode: http.StatusOK, Body: emptyAccountsBody}

	errorBody, err := json.Marshal(test.ServerErrorResponse())
	assert.Nil(t, err)

	errorResponse := &internalhttp.Response{StatusCode: http.StatusInternalServerError, Body: errorBody}

	brokenResponse := &internalhttp.Response{StatusCode: http.StatusInternalServerError, Body: nil}

	api := api.DefaultAPI()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    data.IOutput
		wantErr bool
	}{
		{"GivenAccountsWhenDefaultQueryThenReturnAccounts",
			fields{config: api, request: test.NewRequestMock(accountsResponse, nil)},
			args{urls: &url.Values{}},
			multipleAccounts,
			false},
		{"GivenNoAccountsWhenDefaultQueryThenReturnNilArray",
			fields{config: api, request: test.NewRequestMock(emptyAccountsResponse, nil)},
			args{urls: &url.Values{}},
			&data.AccountsData{},
			false},
		{"WhenNon200ThenReturnErrorMessage",
			fields{config: api, request: test.NewRequestMock(errorResponse, nil)},
			args{urls: &url.Values{}},
			test.ServerErrorResponse(),
			false},
		{"WhenNon200BrokenResponseThenReturnError",
			fields{config: api, request: test.NewRequestMock(brokenResponse, nil)},
			args{urls: &url.Values{}},
			&data.NoOp{},
			true},
		{"WhenHTTPClientThrowsThenReturnError",
			fields{config: api, request: test.NewRequestMock(nil, test.ErrBrokenHTTPClient)},
			args{urls: &url.Values{}},
			&data.NoOp{},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := test.NewTestClient(tt.fields.request)
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
