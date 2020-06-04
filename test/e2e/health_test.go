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

func TestFetch(t *testing.T) {
	client := accountAPIClient.NewAccountAPIClient()
	client.Fetch("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
}

func TestCreate(t *testing.T) {
	client := accountAPIClient.NewAccountAPIClient()
	client.Create(test.NewAccountFromFile("create.json"))
}
