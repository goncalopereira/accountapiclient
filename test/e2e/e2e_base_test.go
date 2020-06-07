package e2e_test

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/data"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	test2 "github.com/goncalopereira/accountapiclient/internal/test"
	client2 "github.com/goncalopereira/accountapiclient/pkg/accountsclient"
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
	Client       *client2.Client
}

//default new client with real http.
func NewClientFromEnv() *client2.Client {
	return client2.NewClient(internalhttp.NewRequest())
}

func (suite *BaseTestSuite) SetupNewAccount(newAccount *data.Data) {
	output, err := suite.Client.Create(newAccount)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &data.Data{}, output)
}

func (suite *BaseTestSuite) NewAccount(id fmt.Stringer) *data.Data {
	newAccount := test2.NewAccountFromFile("create.json")
	newAccount.ID = id.String()

	return newAccount
}

func (suite *BaseTestSuite) SetupTest() {
	suite.setupDeleteAllAccounts(suite.Client)
}

func (suite *BaseTestSuite) SetupSuite() {
	suite.NewAccountID = uuid.New()
	suite.Client = NewClientFromEnv()
}

//try to clean up the api without accessing the DB
//will not work if more than 1 page returned.
func (suite *BaseTestSuite) setupDeleteAllAccounts(client *client2.Client) {
	accounts, err := client.List(&url.Values{})
	assert.Nil(suite.T(), err)

	accountsData := accounts.(*data.AccountsData)

	if accountsData.Accounts == nil {
		return //already empty
	}

	for _, a := range *accountsData.Accounts {
		output, err := client.Delete(a.ID, 0) //versionId does not work on fake api
		assert.Nil(suite.T(), err)
		assert.IsType(suite.T(), &data.NoContent{}, output)
	}
}
