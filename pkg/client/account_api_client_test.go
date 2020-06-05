//nolint:scopelint,funlen
package client_test

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/config"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"github.com/goncalopereira/accountapiclient/pkg/client"
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
	accountBody, _ := json.DataToBody(completeAccount)
	accountResponse := &http.Response{StatusCode: 200, Body: accountBody}

	apiErrorMessage := test.NewErrorMessageFromFile("server-error.json")

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
		{"WhenGivenValidIDThenReturnAccount",
			fields{config: api, request: httptest.NewGetRequestMock(accountResponse, nil)},
			args{id: "1"},
			&data.Output{Account: completeAccount},
			false},
		{"WhenGivenNon200ThenReturnErrorMessage",
			fields{config: api, request: httptest.NewGetRequestMock(errorResponse, nil)},
			args{id: "1"},
			&data.Output{ErrorResponse: apiErrorMessage},
			false},
		{"WhenGivenNon200BrokenResponseThenReturnError",
			fields{config: api, request: httptest.NewGetRequestMock(brokenResponse, nil)},
			args{id: "1"},
			nil,
			true},
		{"WhenHTTPClientThrowsThenReturnError",
			fields{config: api, request: httptest.NewGetRequestMock(nil, fmt.Errorf("boom"))},
			args{id: "1"},
			nil,
			true},
		{"WhenBrokenAPIConfigThrowsThenReturnError",
			fields{config: brokenAPI, request: nil},
			args{id: "1"},
			nil,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := client.NewClient(tt.fields.config, tt.fields.request)

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

func TestClient_Delete(t *testing.T) {
	type fields struct {
		config  config.IAPI
		request http.IRequest
	}
	type args struct {
		id      string
		version int
	}

	deleteResponse := &http.Response{StatusCode: 204}

	apiErrorMessage := test.NewErrorMessageFromFile("server-error.json")

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
		{"WhenGivenValidIDAndVersionThen204Empty",
			fields{config: api, request: httptest.NewDeleteRequestMock(deleteResponse, nil)},
			args{id: "1", version: 1},
			&data.Output{},
			false},
		//includes 404 not found
		//includes 409 specified version incorrect
		{"WhenGivenNon200ThenReturnErrorMessage",
			fields{config: api, request: httptest.NewDeleteRequestMock(errorResponse, nil)},
			args{id: "1", version: 1},
			&data.Output{ErrorResponse: apiErrorMessage},
			false},
		{"WhenGivenNon200BrokenResponseThenReturnError",
			fields{config: api, request: httptest.NewDeleteRequestMock(brokenResponse, nil)},
			args{id: "1", version: 1},
			nil,
			true},
		{"WhenHTTPClientThrowsThenReturnError",
			fields{config: api, request: httptest.NewDeleteRequestMock(nil, fmt.Errorf("boom"))},
			args{id: "1", version: 1},
			nil,
			true},
		{"WhenBrokenAPIConfigThrowsThenReturnError",
			fields{config: brokenAPI, request: nil},
			args{id: "1"},
			nil,
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := client.NewClient(tt.fields.config, tt.fields.request)

			got, err := client.Delete(tt.args.id, tt.args.version)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (err != nil) && tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Delete() got = %v, want %v", got, tt.want)
			}
		})
	}
}
