package json

import (
	"encoding/json"
	"fmt"
)

//json parser and tests for the account/error response marshaling
//will try to hydrate given interface with http response.
func BytesToData(body []byte, data interface{}) error {
	err := json.Unmarshal(body, &data)

	if err != nil {
		return fmt.Errorf("BytesToData: %w", err)
	}

	return nil
}

func DataToBytes(data interface{}) ([]byte, error) {
	bytes, err := json.Marshal(data)

	if err != nil {
		return nil, fmt.Errorf("DataToBytes: %w", err)
	}

	return bytes, nil
}
