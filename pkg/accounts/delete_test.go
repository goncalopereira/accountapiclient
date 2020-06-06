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
	brokenAPI := configtest.NewAPIMock(nil, test.ErrBrokenConfig)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    data.IOutput
		wantErr bool
	}{
		{"WhenGivenValidIDAndVersionThen204Empty",
			fields{config: api, request: httptest.NewGetRequestMock(deleteResponse, nil)},
			args{id: "1", version: 1},
			nil,
			false},
		//includes 404 not found
		//includes 409 specified version incorrect
		{"WhenGivenNon200ThenReturnErrorMessage",
			fields{config: api, request: httptest.NewGetRequestMock(errorResponse, nil)},
			args{id: "1", version: 1},
			apiErrorMessage,
			false},
		{"WhenGivenNon200BrokenResponseThenReturnError",
			fields{config: api, request: httptest.NewGetRequestMock(brokenResponse, nil)},
			args{id: "1", version: 1},
			nil,
			true},
		{"WhenHTTPClientThrowsThenReturnError",
			fields{config: api, request: httptest.NewGetRequestMock(nil, test.ErrBrokenHTTPClient)},
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
			c := accounts.NewClient(tt.fields.config, tt.fields.request)

			got, err := c.Delete(tt.args.id, tt.args.version)
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
