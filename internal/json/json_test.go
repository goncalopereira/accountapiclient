//nolint:scopelint
package json_test

import (
	account "github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"github.com/goncalopereira/accountapiclient/test"
	"testing"
)

func TestRequestBodyToData(t *testing.T) {
	type args struct {
		body []byte
		data interface{}
	}

	var badJSON = args{test.ReadJSON("badjson.txt"), account.Data{}}

	var accountCreate = args{test.ReadJSON("create.json"), account.Data{}}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "BadJSON", args: badJSON, wantErr: true},
		{name: "AccountData", args: accountCreate, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := json.BodyToData(tt.args.body, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("BodyToData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
