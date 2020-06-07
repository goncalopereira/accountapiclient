package main

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/data"
	accountsClient "github.com/goncalopereira/accountapiclient/pkg/accountsclient"
	"github.com/google/uuid"
)

func main() {
	client := accountsClient.NewClient()

	newAccount := data.Account{}
	newAccount.ID = uuid.New()
	//data := &data.Data{Account: newAccount}

	//delete if already exists
	output, err := client.Delete(newAccount.ID, 0)
	if err != nil {
		panic("is API up?")
	}

	fmt.Printf("%+v", output.String())

	/*
		switch t := output.(type) {
		case *data.NoContent:
			fmt.Printf("deleted pre existing account: %s", newAccount.ID)
		case *data.ErrorResponse:
			fmt.Printf("%i", output.(data.ErrorResponse).StatusCode)
			fmt.Printf(output.(data.ErrorResponse).ErrorMessage)
		}
	*/
}
