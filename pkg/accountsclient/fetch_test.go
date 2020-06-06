//nolint:scopelint,funlen
package accountsclient_test

import (
	"encoding/json"
	"github.com/goncalopereira/accountapiclient/internal/api"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/pkg/accountsclient"
	"github.com/goncalopereira/accountapiclient/test"
	httptest "github.com/goncalopereira/accountapiclient/test/http"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestClient_Fetch(t *testing.T) {
	type fields struct {
		config  *api.API
		request http.IRequest
	}

	type args struct {
		id string
	}

	completeAccount := test.NewAccountFromFile("fetch-response.json")
	accountBody, err := json.Marshal(completeAccount)
	assert.Nil(t, err)

	accountResponse := &http.Response{StatusCode: 200, Body: accountBody}

	apiErrorMessage := test.NewErrorMessageFromFile("server-error.json")

	errorBody, err := json.Marshal(apiErrorMessage)
	assert.Nil(t, err)

	errorResponse := &http.Response{StatusCode: 500, Body: errorBody}

	brokenResponse := &http.Response{StatusCode: 500, Body: nil}

	api := api.DefaultAPI()

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    data.IOutput
		wantErr bool
	}{
		{"GivenAccountWhenValidIDThenReturnAccount",
			fields{config: api, request: httptest.NewRequestMock(accountResponse, nil)},
			args{id: "1"},
			completeAccount,
			false},
		{"WhenNon200ThenReturnErrorMessage",
			fields{config: api, request: httptest.NewRequestMock(errorResponse, nil)},
			args{id: "1"},
			apiErrorMessage,
			false},
		{"WhenNon200BrokenResponseThenReturnError",
			fields{config: api, request: httptest.NewRequestMock(brokenResponse, nil)},
			args{id: "1"},
			&data.NoOp{},
			true},
		{"WhenHTTPClientThrowsThenReturnError",
			fields{config: api, request: httptest.NewRequestMock(nil, test.ErrBrokenHTTPClient)},
			args{id: "1"},
			&data.NoOp{},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := accountsclient.NewClient(tt.fields.config, tt.fields.request)

			got, err := client.Fetch(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) && tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fetch() got = %v, want %v", got, tt.want)
			}
		})
	}
}
