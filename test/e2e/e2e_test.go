package e2e

import (
	client2 "github.com/goncalopereira/accountapiclient/pkg/client"
	"github.com/goncalopereira/accountapiclient/test"
	"net/url"
	"testing"
)

//	"github.com/stretchr/testify/assert"

func TestHealth(t *testing.T) {
	client := client2.NewClientFromEnv()
	_, _ = client.Health()
}

var createData = test.NewAccountFromFile("create.json")

func TestBeforeFetch(t *testing.T) {
	client := client2.NewClientFromEnv()
	output, err := client.Fetch(createData.ID)
	if err != nil {
		t.Log(err)
	}
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
	output, _ := client.List(&url.Values{})
	t.Log(output.String())
}

func TestDelete0(t *testing.T) {
	client := client2.NewClientFromEnv()
	output, _ := client.Delete(createData.ID, 0)
	t.Log(output.String())
}

func TestDelete1(t *testing.T) {
	client := client2.NewClientFromEnv()
	output, _ := client.Delete(createData.ID, 1)
	t.Log(output.String())
}
