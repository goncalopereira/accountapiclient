package e2e_test

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/url"
	"testing"
)

//E2E suite for edge cases where no accounts exist
//I prefer to focus on unit tests and prevent E2E
//but was requested to test against the given API
//Ideally I'd use better provider testing instead of this
//as many features are not available against fake api.
type ListTestSuite struct {
	BaseTestSuite
}

func TestListTestSuite(t *testing.T) {
	suite.Run(t, new(ListTestSuite))
}

func (suite *ListTestSuite) SetupTest() {
	suite.BaseTestSuite.SetupTest()

	//add five accounts for pagination tests
	suite.SetupNewAccount(suite.NewAccount(suite.NewAccountID))
	suite.SetupNewAccount(suite.NewAccount(uuid.New()))
	suite.SetupNewAccount(suite.NewAccount(uuid.New()))
	suite.SetupNewAccount(suite.NewAccount(uuid.New()))
	suite.SetupNewAccount(suite.NewAccount(uuid.New()))
}

func (suite *ListTestSuite) TestGivenFiveAccountsWhenListThenListWithFiveAccounts() {
	params := &url.Values{}
	output, err := suite.Client.List(params)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &data.AccountsData{}, output)

	//checking length and not IDs just to keep sorting flexible
	accounts := *output.(*data.AccountsData).Accounts
	assert.Equal(suite.T(), 5, len(accounts))
}

func (suite *ListTestSuite) TestGivenFiveAccountsWhenListWithPage0Size3ThenListWithThreeAccounts() {
	params := &url.Values{}
	params.Add("page[number]", "0")
	params.Add("page[size]", "3")

	output, err := suite.Client.List(params)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &data.AccountsData{}, output)

	accounts := *output.(*data.AccountsData).Accounts
	assert.Equal(suite.T(), 3, len(accounts))
}

func (suite *ListTestSuite) TestGivenFiveAccountsWhenListWithPage1Size3ThenListWithTwoAccounts() {
	params := &url.Values{}
	params.Add("page[number]", "1")
	params.Add("page[size]", "3")

	output, err := suite.Client.List(params)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &data.AccountsData{}, output)

	accounts := *output.(*data.AccountsData).Accounts
	assert.Equal(suite.T(), 2, len(accounts))
}
