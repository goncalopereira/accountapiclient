package data

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/json"
)

type ErrorResponse struct {
	fmt.Stringer
	ErrorMessage string `json:"error_message"`
}

func (e *ErrorResponse) String() string {
	error, _ := json.DataToBody(e)
	return string(error)
}
