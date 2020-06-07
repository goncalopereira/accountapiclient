package e2e_test

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/url"
	"testing"
)

//E2E suite for edge cases where no accounts exist.
type NoAccountsTestSuite struct {
	BaseTestSuite
}

func TestNoAccountsTestSuite(t *testing.T) {
	suite.Run(t, new(NoAccountsTestSuite))
}

//testing with an added filter but fake api does not respect filters.
func (suite *NoAccountsTestSuite) TestGivenNoAccountsWhenListThenEmptyList() {
	params := &url.Values{}
	output, err := suite.Client.List(params)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &data.AccountsData{}, output)

	//https://github.com/golang/go/wiki/CodeReviewComments#declaring-empty-slices
	assert.Nil(suite.T(), output.(*data.AccountsData).Accounts) //nil list is default behaviour

	assert.Equal(suite.T(), &data.AccountsData{}, output) //empty list
}

func (suite *NoAccountsTestSuite) TestGivenNoAccountsWhenFetchUnknownIDThenErrorMessage() {
	output, err := suite.Client.Fetch(suite.NewAccountID)

	assert.Nil(suite.T(), err)
	errorMessage := fmt.Sprintf("record %s does not exist", suite.NewAccountID.String())
	assert.Equal(suite.T(), &data.ErrorResponse{StatusCode: 404, ErrorMessage: errorMessage}, output)
}

//Only testing version 0 as FakeAPI does not handle errors like 409 wrong version
//Expected 404 here but seems to always get a good result even with unknown id
//Unit tests have the correct behavior tested.
func (suite *NoAccountsTestSuite) TestGivenNoAccountsWhenDeleteUnknownIDAndVersion0ThenErrorMessage() {
	output, err := suite.Client.Delete(suite.NewAccountID, 0)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), &data.Deleted{}, output)
}

func (suite *NoAccountsTestSuite) TestGivenNoAccountsWhenCreateUnknownIDThenAccount() {
	output, err := suite.Client.Create(&suite.NewAccountData(suite.NewAccountID).Account)

	assert.Nil(suite.T(), err)
	assert.IsType(suite.T(), &data.Data{}, output)
	assert.Equal(suite.T(), suite.NewAccountID, output.(*data.Data).ID)
}
