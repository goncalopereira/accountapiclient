package data

import (
	"fmt"
	"github.com/goncalopereira/accountapiclient/internal/json"
)

type ErrorResponse struct {
	IOutput      `json:",omitempty"`
	fmt.Stringer `json:",omitempty"`
	ErrorMessage string `json:"error_message"`
	StatusCode   int
}

func (e *ErrorResponse) String() string {
	errorResponse, err := json.DataToBytes(e)
	if err != nil {
		return ""
	}
	return string(errorResponse)
}
