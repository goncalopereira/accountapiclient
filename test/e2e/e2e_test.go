package e2e_test

import (
	"github.com/goncalopereira/accountapiclient/internal/config"
	internalhttp "github.com/goncalopereira/accountapiclient/internal/http"
	client2 "github.com/goncalopereira/accountapiclient/pkg/accounts"
	"github.com/goncalopereira/accountapiclient/test"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

//default new client with real http.
func NewClientFromEnv() *client2.Client {
	return client2.NewClient(config.DefaultAPI(), internalhttp.NewRequest())
}

func TestBeforeFetch(t *testing.T) {
	client := NewClientFromEnv()
	output, _ := client.Fetch(test.NewAccountFromFile("create.json").ID)
	t.Log(output.String())
}

func TestCreate(t *testing.T) {
	client := NewClientFromEnv()
	output, _ := client.Create(test.NewAccountFromFile("create.json"))
	t.Log(output.String())
}

func TestAfterFetch(t *testing.T) {
	client := NewClientFromEnv()
	output, _ := client.Fetch(test.NewAccountFromFile("create.json").ID)
	t.Log(output.String())
}

func TestList(t *testing.T) {
	client := NewClientFromEnv()
	output, err := client.List(&url.Values{})
	assert.Nil(t, err)
	t.Log(output.String())
}

func TestListWithPage2Empty(t *testing.T) {
	client := NewClientFromEnv()
	params := &url.Values{}
	params.Add("page[number]", "1")
	output, err := client.List(params)
	assert.Nil(t, err)
	t.Log(output.String())
}

func TestDelete(t *testing.T) {
	client := NewClientFromEnv()
	output, err := client.Delete(test.NewAccountFromFile("create.json").ID, 0)
	assert.Nil(t, err)
	t.Log(output.String())
}
