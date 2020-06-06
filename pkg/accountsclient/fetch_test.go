//nolint:scopelint,funlen
package accountsclient_test

import (
	"encoding/json"
	"github.com/goncalopereira/accountapiclient/internal/data"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/pkg/accountsclient"
	"github.com/goncalopereira/accountapiclient/test"
	httptest "github.com/goncalopereira/accountapiclient/test/http"
	"github.com/stretchr/testify/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestClient_Fetch(t *testing.T) {
	type fields struct {
		request internalhttp.IRequest
	}

	type args struct {
		id string
	}

	completeAccount := test.NewAccountFromFile("fetch-response.json")
	accountBody, err := json.Marshal(completeAccount)
	assert.Nil(t, err)

	accountResponse := &internalhttp.Response{StatusCode: http.StatusOK, Body: accountBody}

	errorBody, err := json.Marshal(test.ServerErrorResponse())
	assert.Nil(t, err)

	errorResponse := &internalhttp.Response{StatusCode: http.StatusInternalServerError, Body: errorBody}

	brokenResponse := &internalhttp.Response{StatusCode: http.StatusInternalServerError, Body: nil}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    data.IOutput
		wantErr bool
	}{
		{"GivenAccountWhenValidIDThenReturnAccount",
			fields{request: httptest.NewRequestMock(accountResponse, nil)},
			args{id: "1"},
			completeAccount,
			false},
		{"WhenNon200ThenReturnErrorMessage",
			fields{request: httptest.NewRequestMock(errorResponse, nil)},
			args{id: "1"},
			test.ServerErrorResponse(),
			false},
		{"WhenNon200BrokenResponseThenReturnError",
			fields{request: httptest.NewRequestMock(brokenResponse, nil)},
			args{id: "1"},
			&data.NoOp{},
			true},
		{"WhenHTTPClientThrowsThenReturnError",
			fields{request: httptest.NewRequestMock(nil, test.ErrBrokenHTTPClient)},
			args{id: "1"},
			&data.NoOp{},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := accountsclient.NewClient(tt.fields.request)

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
