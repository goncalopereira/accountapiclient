package e2e_test

import (
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/goncalopereira/accountapiclient/internal/data/account"
	"github.com/goncalopereira/accountapiclient/test"
	"github.com/goncalopereira/accountapiclient/test/e2e"
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
type OneAccountTestSuite struct {
	e2e.BaseTestSuite
}

func TestOneAccountTestSuite(t *testing.T) {
	suite.Run(t, new(OneAccountTestSuite))
}

func (suite *OneAccountTestSuite) SetupTest() {
	suite.BaseTestSuite.SetupTest()
	suite.SetupNewAccount(suite.NewAccount(suite.NewAccountID))
}

func (suite *OneAccountTestSuite) TestGivenOneAccountWhenListThenListWithOneAccount() {
	params := &url.Values{}
	output, err := suite.Client.List(params)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &account.AccountsData{}, output)

	accounts := *output.(*account.AccountsData).Accounts
	firstAccount := accounts[0]
	assert.EqualValues(suite.T(), suite.NewAccountID.String(), firstAccount.ID)
}

func (suite *OneAccountTestSuite) TestGivenOneAccountWhenFetchIDThenAccount() {
	output, err := suite.Client.Fetch(suite.NewAccountID.String())

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &account.Data{}, output)
	assert.Equal(suite.T(), suite.NewAccountID.String(), output.(*account.Data).ID)
}

//Only testing version 0 as FakeAPI does not handle errors like 409 wrong version
//Expected 404 here but seems to always get a good result even with unknown id
//Unit tests have the correct behavior tested.
func (suite *OneAccountTestSuite) TestGivenOneAccountWhenDeleteIDAndVersion0ThenNoContent() {
	output, err := suite.Client.Delete(suite.NewAccountID.String(), 0)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), &data.NoContent{}, output)
}

func (suite *OneAccountTestSuite) TestGivenOneAccountWhenCreateSameIDThenErrorMessage() {
	output, err := suite.Client.Create(suite.NewAccount(suite.NewAccountID))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), test.DuplicateAccountErrorResponse(), output.(*data.ErrorResponse))
}
