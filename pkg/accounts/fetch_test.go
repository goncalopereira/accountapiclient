//nolint:scopelint,funlen
package accounts_test

import (
	"github.com/goncalopereira/accountapiclient/internal/config"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"github.com/goncalopereira/accountapiclient/pkg/accounts"
	"github.com/goncalopereira/accountapiclient/test"
	configtest "github.com/goncalopereira/accountapiclient/test/config"
	httptest "github.com/goncalopereira/accountapiclient/test/http"
	"reflect"
	"testing"
)

func TestClient_Fetch(t *testing.T) {
	type fields struct {
		config  config.IAPI
		request http.IRequest
	}

	type args struct {
		id string
	}

	completeAccount := test.NewAccountFromFile("complete-account.json")
	accountBody, _ := json.DataToBytes(completeAccount)
	accountResponse := &http.Response{StatusCode: 200, Body: accountBody}

	apiErrorMessage := test.NewErrorMessageFromFile("server-error.json")

	errorBody, _ := json.DataToBytes(apiErrorMessage)
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
		{"GivenAccountWhenValidIDThenReturnAccount",
			fields{config: api, request: httptest.NewGetRequestMock(accountResponse, nil)},
			args{id: "1"},
			completeAccount,
			false},
		{"WhenNon200ThenReturnErrorMessage",
			fields{config: api, request: httptest.NewGetRequestMock(errorResponse, nil)},
			args{id: "1"},
			apiErrorMessage,
			false},
		{"WhenNon200BrokenResponseThenReturnError",
			fields{config: api, request: httptest.NewGetRequestMock(brokenResponse, nil)},
			args{id: "1"},
			&data.NoOp{},
			true},
		{"WhenHTTPClientThrowsThenReturnError",
			fields{config: api, request: httptest.NewGetRequestMock(nil, test.ErrBrokenHTTPClient)},
			args{id: "1"},
			&data.NoOp{},
			true},
		{"WhenBrokenAPIConfigThrowsThenReturnError",
			fields{config: brokenAPI, request: nil},
			args{id: "1"},
			&data.NoOp{},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := accounts.NewClient(tt.fields.config, tt.fields.request)

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
