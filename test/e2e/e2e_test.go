package e2e_test

import (
	client2 "github.com/goncalopereira/accountapiclient/pkg/client"
	"github.com/goncalopereira/accountapiclient/test"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

var createData = test.NewAccountFromFile("create.json")

func TestBeforeFetch(t *testing.T) {
	client := client2.NewClientFromEnv()
	output, _ := client.Fetch(createData.ID)
	t.Log(output.String())
}

func TestCreate(t *testing.T) {
	client := client2.NewClientFromEnv()
	output, _ := client.Create(createData)
	t.Log(output.String())
}

func TestAfterFetch(t *testing.T) {
	client := client2.NewClientFromEnv()
	output, _ := client.Fetch(createData.ID)
	t.Log(output.String())
}

func TestList(t *testing.T) {
	client := client2.NewClientFromEnv()
	output, err := client.List(&url.Values{})
	assert.Nil(t, err)
	t.Log(output.String())
}

func TestListWithPage2Empty(t *testing.T) {
	client := client2.NewClientFromEnv()
	params := &url.Values{}
	params.Add("page[number]", "1")
	output, err := client.List(params)
	assert.Nil(t, err)
	t.Log(output.String())
}

/*func TestDelete1(t *testing.T) {
	client := client2.NewClientFromEnv()
	output, err := client.Delete(createData.ID, 0)
	assert.Nil(t, output.ErrorResponse)
	assert.Nil(t, err)
	t.Log(output.String())
}
*/
