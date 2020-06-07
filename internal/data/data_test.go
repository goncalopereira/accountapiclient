package data_test

import (
	test2 "github.com/goncalopereira/accountapiclient/internal/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestData_WhenReadingAccountFromFileThenReturnAccount(t *testing.T) {
	data := test2.NewAccountFromFile("fetch-response.json")

	accountID, err := uuid.Parse("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	assert.Nil(t, err)

	orgID, err := uuid.Parse("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	assert.Nil(t, err)

	assert.Equal(t, accountID, data.ID)
	assert.Equal(t, orgID, data.OrganisationID)
	assert.Equal(t, "GB", data.Country)
	assert.Equal(t, "GB", data.Attributes.Country)

	assert.Equal(t, []string{"Sam Holder"}, data.AlternativeNames) //test array
	assert.Equal(t, []string{"Samantha Holder"}, data.Name)

	assert.Equal(t, "2017-07-23", data.PrivateIdentificationBirthDate)                   //private ID
	assert.Equal(t, []string{"10 Avenue des Champs"}, data.PrivateIdentificationAddress) //private ID

	assert.Equal(t, "123654", data.OrganisationIdentificationIdentification) //org ID

	assert.Equal(t, "1970-01-01", data.OrganisationIdentification.Actors[0].BirthDate) //Actors
	assert.Equal(t, "Jeff Page", data.OrganisationIdentification.Actors[0].Names[0])   //Actors

	masterID, err := uuid.Parse("a52d13a4-f435-4c00-cfad-f5e7ac5972df")
	assert.Nil(t, err)

	eventsID, err := uuid.Parse("c1023677-70ee-417a-9a6a-e211241f1e9c")
	assert.Nil(t, err)

	assert.Equal(t, masterID, data.RelationshipsMasterAccount.Data[0].ID) //master account
	assert.Equal(t, eventsID, data.RelationshipsAccountEvents.Data[0].ID) //account events
}
