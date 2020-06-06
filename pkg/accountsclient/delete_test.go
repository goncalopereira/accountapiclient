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

func TestClient_Delete(t *testing.T) {
	type fields struct {
		request internalhttp.IRequest
	}

	type args struct {
		id      string
		version int
	}

	deleteResponse := &internalhttp.Response{StatusCode: 204}

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
		{"WhenGivenValidIDAndVersionThen204Empty",
			fields{request: httptest.NewRequestMock(deleteResponse, nil)},
			args{id: "1", version: 1},
			&data.NoContent{},
			false},
		//includes 404 not found
		//includes 409 specified version incorrect
		{"WhenGivenNon200ThenReturnErrorMessage",
			fields{request: httptest.NewRequestMock(errorResponse, nil)},
			args{id: "1", version: 1},
			test.ServerErrorResponse(),
			false},
		{"WhenGivenNon200BrokenResponseThenReturnError",
			fields{request: httptest.NewRequestMock(brokenResponse, nil)},
			args{id: "1", version: 1},
			&data.NoOp{},
			true},
		{"WhenHTTPClientThrowsThenReturnError",
			fields{request: httptest.NewRequestMock(nil, test.ErrBrokenHTTPClient)},
			args{id: "1", version: 1},
			&data.NoOp{},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := accountsclient.NewClient(tt.fields.request)

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
