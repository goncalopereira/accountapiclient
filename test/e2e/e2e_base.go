package e2e

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	client2 "github.com/goncalopereira/accountapiclient/pkg/accountsclient"
	"github.com/goncalopereira/accountapiclient/test"
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

func (suite *BaseTestSuite) SetupNewAccount(newAccount *account.Data) {
	output, err := suite.Client.Create(newAccount)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &account.Data{}, output)
}

func (suite *BaseTestSuite) NewAccount(id fmt.Stringer) *account.Data {
	newAccount := test.NewAccountFromFile("create.json")
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

	accountsData := accounts.(*account.AccountsData)

	if accountsData.Accounts == nil {
		return //already empty
	}

	for _, a := range *accountsData.Accounts {
		output, err := client.Delete(a.ID, 0) //versionId does not work on fake api
		assert.Nil(suite.T(), err)
		assert.IsType(suite.T(), &data.NoContent{}, output)
	}
}
