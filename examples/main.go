package main

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/data"
	accountsClient "github.com/goncalopereira/accountapiclient/pkg/accountsclient"
	"github.com/google/uuid"
	"net/url"
)

func main() {
	client := accountsClient.NewClient()

	newAccount := accountsClient.NewAccount(uuid.New(), "Gb")
	data := &data.Data{Account: *newAccount}

	newAccountData, err := client.Create(data)
	if err != nil {
		panic("is API up?")
	}

	fmt.Printf("create %+w\n", newAccountData)

	fetchedData, err := client.Fetch(newAccount.ID)
	if err != nil {
		panic("is API up?")
	}

	fmt.Printf("fetch %+w\n", fetchedData)

	accountsData, err := client.List(&url.Values{})
	if err != nil {
		panic("is API up?")
	}

	fmt.Printf("list %+w\n", accountsData)

	deleted, err := client.Delete(newAccount.ID, 0)
	if err != nil {
		panic("is API up?")
	}

	fmt.Printf("delete %+w\n", deleted)
}
