//nolint:scopelint,funlen
package accountsclient_test

import (
	"encoding/json"
	"github.com/goncalopereira/accountapiclient/internal/data"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	"github.com/goncalopereira/accountapiclient/internal/test"
	"github.com/google/uuid"
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
		id      uuid.UUID
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
			fields{request: test.NewRequestMock(deleteResponse, nil)},
			args{id: uuid.New(), version: 1},
			&data.NoContent{},
			false},
		//includes 404 not found
		//includes 409 specified version incorrect
		{"WhenGivenNon200ThenReturnErrorMessage",
			fields{request: test.NewRequestMock(errorResponse, nil)},
			args{id: uuid.New(), version: 1},
			test.ServerErrorResponse(),
			false},
		{"WhenGivenNon200BrokenResponseThenReturnError",
			fields{request: test.NewRequestMock(brokenResponse, nil)},
			args{id: uuid.New(), version: 1},
			&data.NoOp{},
			true},
		{"WhenHTTPClientThrowsThenReturnError",
			fields{request: test.NewRequestMock(nil, test.ErrBrokenHTTPClient)},
			args{id: uuid.New(), version: 1},
			&data.NoOp{},
			true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := test.NewTestClient(tt.fields.request)

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
