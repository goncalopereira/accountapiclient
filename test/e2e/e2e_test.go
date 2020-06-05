package e2e

import (
	accountAPIClient "github.com/goncalopereira/accountapiclient/pkg"
	"github.com/goncalopereira/accountapiclient/test"
	"net/url"
	"testing"
)

//	"github.com/stretchr/testify/assert"

func TestHealth(t *testing.T) {
	client := accountAPIClient.NewDefaultAccountAPIClient()
	_, _ = client.Health()
}

var createData = test.NewAccountFromFile("create.json")

func TestBeforeFetch(t *testing.T) {
	client := accountAPIClient.NewDefaultAccountAPIClient()
	_, _ = client.Fetch(createData.ID)
}

func TestCreate(t *testing.T) {
	client := accountAPIClient.NewDefaultAccountAPIClient()
	_, _ = client.Create(createData)
}

func TestAfterFetch(t *testing.T) {
	client := accountAPIClient.NewDefaultAccountAPIClient()
	_, _ = client.Fetch(createData.ID)
}

func TestList(t *testing.T) {
	client := accountAPIClient.NewDefaultAccountAPIClient()
	_, _ = client.List(&url.Values{})
}

func TestDelete0(t *testing.T) {
	client := accountAPIClient.NewDefaultAccountAPIClient()
	_, _ = client.Delete(createData.ID, 0)
}

func TestDelete1(t *testing.T) {
	client := accountAPIClient.NewDefaultAccountAPIClient()
	_, _ = client.Delete(createData.ID, 1)
}
