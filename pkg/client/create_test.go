//nolint:scopelint,funlen
package client_test

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/config"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"github.com/goncalopereira/accountapiclient/pkg/client"
	"github.com/goncalopereira/accountapiclient/test"
	configtest "github.com/goncalopereira/accountapiclient/test/config"
	httptest "github.com/goncalopereira/accountapiclient/test/http"
	"reflect"
	"testing"
)

func TestClient_Create(t *testing.T) {
	type fields struct {
		config  config.IAPI
		request http.IRequest
	}
	type args struct {
		account *account.Account
	}

	createdAccount := test.NewAccountFromFile("create-response.json")
	accountBody, _ := json.DataToBody(createdAccount)
	accountResponse := &http.Response{StatusCode: 201, Body: accountBody}

	apiErrorMessage := test.DuplicateAccountErrorResponse()

	errorBody, _ := json.DataToBody(apiErrorMessage)
	errorResponse := &http.Response{StatusCode: 500, Body: errorBody}

	brokenResponse := &http.Response{StatusCode: 500, Body: nil}

	api := config.DefaultAPI()
	brokenAPI := configtest.NewAPIMock(nil, fmt.Errorf("broken config"))

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *data.Output
		wantErr bool
	}{
		{"GivenNoAccountWhenPostAccountThenReturnAccount",
			fields{config: api, request: httptest.NewPostRequestMock(accountResponse, nil)},
			args{account: test.AccountCreateRequest()},
			&data.Output{Account: createdAccount},
			false},
		{name: "GivenAccountWhenPostSameIDThenReturnErrorMessage", //409 conflict existing
			fields: fields{config: api, request: httptest.NewPostRequestMock(errorResponse, nil)},
			args:   args{account: test.AccountCreateRequest()},
			want:   &data.Output{ErrorResponse: test.DuplicateAccountErrorResponse()}},
		{"WhenGivenNon200BrokenResponseThenReturnError",
			fields{config: api, request: httptest.NewPostRequestMock(brokenResponse, nil)},
			args{account: test.AccountCreateRequest()},
			nil,
			true},
		/*		{"WhenHTTPClientThrowsThenReturnError",
				fields{config: api, request: httptest.NewGetRequestMock(nil, fmt.Errorf("boom"))},
				args{account: test.AccountCreateRequest()},
				nil,
				true},*/
		{"WhenBrokenAPIConfigThrowsThenReturnError",
			fields{config: brokenAPI, request: nil},
			args{account: test.AccountCreateRequest()},
			nil,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := client.NewClient(tt.fields.config, tt.fields.request)

			got, err := client.Create(tt.args.account)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) && tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}
