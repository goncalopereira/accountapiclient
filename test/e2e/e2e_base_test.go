package e2e_test

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/test"
	"github.com/goncalopereira/accountapiclient/pkg/accountsclient"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/url"
)

//I prefer to focus on unit tests and prevent E2E
//but was requested to test against the given API
//Ideally I'd use better provider testing instead of this
//as many features are not available against fake api.
type BaseTestSuite struct {
	suite.Suite
	NewAccountID uuid.UUID
	Client       *accountsclient.Client
}

func (suite *BaseTestSuite) SetupNewAccount(newAccount *data.Data) {
	output, err := suite.Client.Create(&newAccount.Account)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &data.Data{}, output)
}

func (suite *BaseTestSuite) NewAccountData(id uuid.UUID) *data.Data {
	newAccount := test.NewAccountDataFromFile("create-request.json")
	newAccount.ID = id

	return newAccount
}

func (suite *BaseTestSuite) SetupTest() {
	suite.setupDeleteAllAccounts(suite.Client)
}

func (suite *BaseTestSuite) SetupSuite() {
	suite.NewAccountID = uuid.New()
	suite.Client = accountsclient.NewClient()
}

//try to clean up the api without accessing the DB
//will not work if more than 1 page returned.
func (suite *BaseTestSuite) setupDeleteAllAccounts(client *accountsclient.Client) {
	accounts, err := client.List(&url.Values{})
	assert.Nil(suite.T(), err)

	accountsData := accounts.(*data.AccountsData)

	if accountsData.Accounts == nil {
		return //already empty
	}

	for _, a := range *accountsData.Accounts {
		output, err := client.Delete(a.ID, 0) //versionId does not work on fake api
		assert.Nil(suite.T(), err)
		assert.IsType(suite.T(), &data.Deleted{}, output)
	}
}
