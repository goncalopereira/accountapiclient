package e2e

import (
	"github.com/goncalopereira/accountapiclient/internal/config"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	client2 "github.com/goncalopereira/accountapiclient/pkg/accountsclient"
	"github.com/goncalopereira/accountapiclient/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"net/url"
	"strconv"
)

//I prefer to focus on unit tests and prevent E2E
//but was requested to test against the given API
//Ideally I'd use better provider testing instead of this
//as many features are not available against fake api.
type BaseTestSuite struct {
	suite.Suite
	NewAccountID uuid.UUID
	Client       *client2.Client
	BankID       string
}

//default new client with real http.
func NewClientFromEnv() *client2.Client {
	return client2.NewClient(config.DefaultAPI(), internalhttp.NewRequest())
}

const MaxBankID = 999999
const MinBankID = 100000

func randomBankID() int {
	return rand.Intn(MaxBankID-MinBankID) - MinBankID
}

func (suite *BaseTestSuite) SetupNewAccount(newAccount *account.Data) {
	suite.T().Logf("creating setup account %s", newAccount.ID)
	output, err := suite.Client.Create(newAccount)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &account.Data{}, output)
}

func (suite *BaseTestSuite) NewAccount() *account.Data {
	newAccount := test.NewAccountFromFile("create.json")
	newAccount.ID = suite.NewAccountID.String()

	return newAccount
}

func (suite *BaseTestSuite) SetupTest() {
	suite.setupDeleteAllAccounts(suite.Client)
}

func (suite *BaseTestSuite) SetupSuite() {
	suite.NewAccountID = uuid.New()
	suite.Client = NewClientFromEnv()
	suite.BankID = strconv.Itoa(randomBankID())
}

//try to clean up the api without accessing the DB
//will not work if more than 1 page returned.
func (suite *BaseTestSuite) setupDeleteAllAccounts(client *client2.Client) {
	accounts, _ := client.List(&url.Values{})
	accountsData := accounts.(*account.AccountsData)

	if accountsData.Accounts == nil {
		return //already empty
	}

	for _, a := range *accountsData.Accounts {
		suite.T().Logf("deleting %s", a.ID)
		output, err := client.Delete(a.ID, 0) //versionId does not work on fake api
		assert.Nil(suite.T(), err)
		assert.IsType(suite.T(), &data.NoContent{}, output)
	}
}
