package e2e_test

import (
	test2 "github.com/goncalopereira/accountapiclient/internal/test"
	"github.com/goncalopereira/accountapiclient/pkg/accountsclient"
	"net/url"
	"testing"

	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

//E2E suite for edge cases where no accounts exist
//I prefer to focus on unit tests and prevent E2E
//but was requested to test against the given API
//Ideally I'd use better provider testing instead of this
//as many features are not available against fake api.
type OneAccountTestSuite struct {
	BaseTestSuite
}

func TestOneAccountTestSuite(t *testing.T) {
	suite.Run(t, new(OneAccountTestSuite))
}

func (suite *OneAccountTestSuite) SetupTest() {
	suite.BaseTestSuite.SetupTest()
	suite.SetupNewAccountData(accountsclient.NewAccount(suite.NewAccountID, accountsclient.GB))
}

func (suite *OneAccountTestSuite) TestGivenOneAccountWhenListThenListWithOneAccount() {
	params := &url.Values{}
	output, err := suite.Client.List(params)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &data.AccountsData{}, output)

	accounts := *output.(*data.AccountsData).Accounts
	firstAccount := accounts[0]
	assert.EqualValues(suite.T(), suite.NewAccountID, firstAccount.ID)
}

func (suite *OneAccountTestSuite) TestGivenOneAccountWhenFetchIDThenAccount() {
	output, err := suite.Client.Fetch(suite.NewAccountID)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &data.Data{}, output)
	assert.Equal(suite.T(), suite.NewAccountID, output.(*data.Data).ID)
}

//Only testing version 0 as FakeAPI does not handle errors like 409 wrong version
//Expected 404 here but seems to always get a good result even with unknown id
//Unit tests have the correct behavior tested.
func (suite *OneAccountTestSuite) TestGivenOneAccountWhenDeleteIDAndVersion0ThenNoContent() {
	output, err := suite.Client.Delete(suite.NewAccountID, 0)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), &data.Deleted{}, output)
}

func (suite *OneAccountTestSuite) TestGivenOneAccountWhenCreateSameIDThenErrorMessage() {
	output, err := suite.Client.Create(accountsclient.NewAccount(suite.NewAccountID, accountsclient.GB))

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), test2.DuplicateAccountErrorResponse(), output.(*data.ErrorResponse))
}
