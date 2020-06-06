package e2e_test

import (
	"fmt"
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
	"testing"
)

//E2E suite for edge cases where no accounts exist
//I prefer to focus on unit tests and prevent E2E
//but was requested to test against the given API
//Ideally I'd use better provider testing instead of this
//as many features are not available against fake api.
type NoAccountsTestSuite struct {
	suite.Suite
	NewAccountID uuid.UUID
	client       *client2.Client
	BankID       string
}

//default new client with real http.
func NewClientFromEnv() *client2.Client {
	return client2.NewClient(config.DefaultAPI(), internalhttp.NewRequest())
}

func randomBankID() int {
	return rand.Intn(999999-100000) - 100000
}

func TestNoAccountsTestSuite(t *testing.T) {
	suite.Run(t, new(NoAccountsTestSuite))
}

//try to clean up the api without accessing the DB
//will not work if more than 1 page returned.
func setupDeleteAllAccounts(client *client2.Client) {
	accounts, _ := client.List(&url.Values{})
	accountsData := accounts.(*account.AccountsData)

	if accountsData.Accounts == nil {
		return //already empty
	}

	for _, a := range *accountsData.Accounts {
		_, _ = client.Delete(a.ID, 0) //versionId does not work on fake api
	}
}

func (suite *NoAccountsTestSuite) SetupTest() {
	setupDeleteAllAccounts(suite.client)
}

func (suite *NoAccountsTestSuite) SetupSuite() {
	suite.NewAccountID = uuid.New()
	suite.client = NewClientFromEnv()
	suite.BankID = strconv.Itoa(randomBankID())
}

//testing with an added filter but fake api does not respect filters.
func (suite *NoAccountsTestSuite) TestGivenNoAccountsWhenListThenEmptyList() {
	params := &url.Values{}
	params.Add("filter[bank_id]", suite.BankID)
	output, err := suite.client.List(params)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &account.AccountsData{}, output)

	//https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
	assert.Nil(suite.T(), output.(*account.AccountsData).Accounts) //nil list is default behaviour

	assert.Equal(suite.T(), &account.AccountsData{}, output) //empty list
}

func (suite *NoAccountsTestSuite) TestGivenNoAccountsWhenFetchUnknownIDThenErrorMessage() {
	output, err := suite.client.Fetch(suite.NewAccountID.String())

	assert.Nil(suite.T(), err)
	errorMessage := fmt.Sprintf("record %s does not exist", suite.NewAccountID.String())
	assert.Equal(suite.T(), &data.ErrorResponse{StatusCode: 404, ErrorMessage: errorMessage}, output)
}

//Only testing version 0 as FakeAPI does not handle errors like 409 wrong version
//Expected 404 here but seems to always get a good result even with unknown id
//Unit tests have the correct behavior tested.
func (suite *NoAccountsTestSuite) TestGivenNoAccountsWhenDeleteUnknownIDAndVersion0ThenErrorMessage() {
	output, err := suite.client.Delete(suite.NewAccountID.String(), 0)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), &data.NoContent{}, output)
}

func (suite *NoAccountsTestSuite) TestGivenNoAccountsWhenCreateUnknownIDThenAccount() {
	newAccount := test.NewAccountFromFile("create.json")
	newAccount.ID = suite.NewAccountID.String()
	output, err := suite.client.Create(newAccount)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &account.Data{}, output)
	assert.Equal(suite.T(), suite.NewAccountID.String(), output.(*account.Data).ID)
}
