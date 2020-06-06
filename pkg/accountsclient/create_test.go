//nolint:scopelint,funlen
package accountsclient_test

import (
	"encoding/json"
	"github.com/goncalopereira/accountapiclient/internal/api"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/pkg/accountsclient"
	"github.com/goncalopereira/accountapiclient/test"
	httptest "github.com/goncalopereira/accountapiclient/test/http"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestClient_Create(t *testing.T) {
	type fields struct {
		config  *api.API
		request http.IRequest
	}

	type args struct {
		account *account.Data
	}

	createdAccount := test.NewAccountFromFile("create-response.json")

	accountBody, err := json.Marshal(createdAccount)
	assert.Nil(t, err)

	accountResponse := &http.Response{StatusCode: 201, Body: accountBody}

	apiErrorMessage := test.DuplicateAccountErrorResponse()

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
		{"GivenNoAccountWhenPostAccountThenReturnAccount",
			fields{config: api, request: httptest.NewRequestMock(accountResponse, nil)},
			args{account: test.AccountCreateRequest()},
			createdAccount,
			false},
		{name: "GivenAccountWhenPostSameIDThenReturnErrorMessage", //409 conflict existing
			fields: fields{config: api, request: httptest.NewRequestMock(errorResponse, nil)},
			args:   args{account: test.AccountCreateRequest()},
			want:   test.DuplicateAccountErrorResponse()},
		{"WhenGivenNon200BrokenResponseThenReturnError",
			fields{config: api, request: httptest.NewRequestMock(brokenResponse, nil)},
			args{account: test.AccountCreateRequest()},
			&data.NoOp{},
			true},
		{"WhenHTTPClientThrowsThenReturnError",
			fields{config: api, request: httptest.NewRequestMock(nil, test.ErrBrokenHTTPClient)},
			args{account: test.AccountCreateRequest()},
			&data.NoOp{},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := accountsclient.NewClient(tt.fields.config, tt.fields.request)

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
