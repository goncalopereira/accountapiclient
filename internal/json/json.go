package json

import (
	"encoding/json"
	"fmt"
)

//json parser and tests for the account/error response marshaling
//will try to hydrate given interface with http response
func BodyToData(body []byte, data interface{}) error {
	unmarshallErr := json.Unmarshal(body, &data)

	if unmarshallErr != nil {
		return fmt.Errorf("BodyToData: %w", unmarshallErr)
	}

	return nil
}

func DataToBody(data interface{}) ([]byte, error) {
	bytes, marshalErr := json.Marshal(data)

	if marshalErr != nil {
		return nil, fmt.Errorf("DataToBody: %w", marshalErr)
	}

	return bytes, nil
}
