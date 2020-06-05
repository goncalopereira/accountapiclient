//nolint:dogsled
package e2e

import (
	http "github.com/goncalopereira/accountapiclient/internal/commands"
	accountAPIClient "github.com/goncalopereira/accountapiclient/pkg"
	"github.com/goncalopereira/accountapiclient/test"
	"testing"
)

//	"github.com/stretchr/testify/assert"

func TestHealth(t *testing.T) {
	err := http.Health()

	if err != nil {
		t.Errorf("Healthcheck failed with error: %v", err)
	}
}

var createData = test.NewAccountFromFile("create.json")

func TestBeforeFetch(t *testing.T) {
	client := accountAPIClient.NewAccountAPIClient()
	_, _, _ = client.Fetch(createData.ID)
}

func TestCreate(t *testing.T) {
	client := accountAPIClient.NewAccountAPIClient()
	_, _, _ = client.Create(createData)
}

func TestAfterFetch(t *testing.T) {
	client := accountAPIClient.NewAccountAPIClient()
	_, _, _ = client.Fetch(createData.ID)
}

func TestList(t *testing.T) {
	client := accountAPIClient.NewAccountAPIClient()
	_, _, _ = client.List(map[string]string{})
}

func TestDelete0(t *testing.T) {
	client := accountAPIClient.NewAccountAPIClient()
	_, _ = client.Delete(createData.ID, 0)
}

func TestDelete1(t *testing.T) {
	client := accountAPIClient.NewAccountAPIClient()
	_, _ = client.Delete(createData.ID, 1)
}
