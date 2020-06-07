//nolint:scopelint,funlen
package accountsclient_test

import (
	"encoding/json"
	"github.com/goncalopereira/accountapiclient/internal/data"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/internal/test"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//createRequestAccount returns valid static Data with an Account.
func createRequestAccount() *data.Account {
	account := test.NewAccountDataFromFile("create-request.json")
	return &account.Account
}

func TestClient_Create(t *testing.T) {
	type fields struct {
		request internalhttp.IRequest
	}

	type args struct {
		account *data.Account
	}

	createdAccount := test.NewAccountDataFromFile("create-response.json")

	accountBody, err := json.Marshal(createdAccount)
	assert.Nil(t, err)

	accountResponse := &internalhttp.Response{StatusCode: 201, Body: accountBody}

	apiErrorMessage := test.DuplicateAccountErrorResponse()

	errorBody, err := json.Marshal(apiErrorMessage)
	assert.Nil(t, err)

	errorResponse := &internalhttp.Response{StatusCode: 500, Body: errorBody}

	brokenResponse := &internalhttp.Response{StatusCode: 500, Body: nil}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    data.IOutput
		wantErr bool
	}{
		{"GivenNoAccountWhenPostAccountThenReturnAccount",
			fields{request: test.NewRequestMock(accountResponse, nil)},
			args{account: createRequestAccount()},
			createdAccount,
			false},
		{name: "GivenAccountWhenPostSameIDThenReturnErrorMessage", //409 conflict existing
			fields: fields{request: test.NewRequestMock(errorResponse, nil)},
			args:   args{account: createRequestAccount()},
			want:   test.DuplicateAccountErrorResponse()},
		{"WhenGivenNon200BrokenResponseThenReturnError",
			fields{request: test.NewRequestMock(brokenResponse, nil)},
			args{account: createRequestAccount()},
			&data.NoOp{},
			true},
		{"WhenHTTPClientThrowsThenReturnError",
			fields{request: test.NewRequestMock(nil, test.ErrBrokenHTTPClient)},
			args{account: createRequestAccount()},
			&data.NoOp{},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := test.NewTestClient(tt.fields.request)

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
