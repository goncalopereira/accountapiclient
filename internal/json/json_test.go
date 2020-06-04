//nolint:scopelint,lll
package json_test

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	account "github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/internal/json"
	"github.com/goncalopereira/accountapiclient/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestBodyToData(t *testing.T) {
	type args struct {
		body []byte
		data interface{}
	}

	var healthUp = args{test.ReadJSON("health-up.json"), data.Health{}}
	var badJSON = args{test.ReadJSON("badjson.txt"), data.Health{}}
	var accountCreate = args{test.ReadJSON("create.json"), account.NewEmptyAccount()}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "Health", args: healthUp, wantErr: false},
		{name: "BadJSON", args: badJSON, wantErr: true},
		{name: "Account", args: accountCreate, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := json.BodyToData(tt.args.body, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("BodyToData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataToBody_WhenValidAccountThenTransformToJSONAndBackAndReturnAccount(t *testing.T) {
	attributes := account.Attributes{Country: "GB"}
	attributes.PrivateIdentification = account.PrivateIdentification{PrivateIdentificationBirthCountry: "PT"}
	attributes.OrganisationIdentification = account.OrganisationIdentification{OrganisationIdentificationCountry: "US"}
	originalAccount := account.NewAccount("newid", "neworganisationId", attributes)

	result, _ := json.DataToBody(originalAccount)

	newAccount := account.NewEmptyAccount()
	_ = json.BodyToData(result, &newAccount)

	assert.Equal(t, "newid", newAccount.LinkData.ID)
	assert.Equal(t, "neworganisationId", newAccount.OrganisationID)
	assert.Equal(t, "PT", newAccount.Attributes.PrivateIdentificationBirthCountry)
	assert.Equal(t, "US", newAccount.Attributes.OrganisationIdentificationCountry)
	assert.Equal(t, "GB", newAccount.Attributes.Country)
}

func TestDataToBody_WhenReadingAccountFromFileThenReturnAccount(t *testing.T) {
	newAccount := test.NewAccountFromFile("complete-account.json")

	assert.Equal(t, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", newAccount.ID)
	assert.Equal(t, "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c", newAccount.OrganisationID)
	assert.Equal(t, "GB", newAccount.Country)
	assert.Equal(t, "GB", newAccount.Attributes.Country)

	assert.Equal(t, []string{"Sam Holder"}, newAccount.AlternativeNames) //test array
	assert.Equal(t, []string{"Samantha Holder"}, newAccount.Name)

	assert.Equal(t, "2017-07-23", newAccount.PrivateIdentificationBirthDate)                   //private ID
	assert.Equal(t, []string{"10 Avenue des Champs"}, newAccount.PrivateIdentificationAddress) //private ID

	assert.Equal(t, "123654", newAccount.OrganisationIdentificationIdentification) //org ID

	assert.Equal(t, "1970-01-01", newAccount.OrganisationIdentification.Actors[0].BirthDate) //Actors
	assert.Equal(t, "Jeff Page", newAccount.OrganisationIdentification.Actors[0].Names[0])   //Actors

	assert.Equal(t, "a52d13a4-f435-4c00-cfad-f5e7ac5972df", newAccount.RelationshipsMasterAccount.Data[0].ID) //master account
	assert.Equal(t, "c1023677-70ee-417a-9a6a-e211241f1e9c", newAccount.RelationshipsAccountEvents.Data[0].ID) //account events
}
