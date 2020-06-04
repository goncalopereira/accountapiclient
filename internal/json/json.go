package json

import (
	"encoding/json"
	"fmt"
)

func BodyToData(body []byte, data interface{}) error {
	unmarshallErr := json.Unmarshal(body, &data)

	if unmarshallErr != nil {
		return fmt.Errorf("BodyToData: %v", unmarshallErr.Error())
	}

	return nil
}

func DataToBody(data interface{}) ([]byte, error) {
	bytes, marshalErr := json.Marshal(data)

	if marshalErr != nil {
		return nil, fmt.Errorf("DataToBody: %v", marshalErr.Error())
	}

	return bytes, nil
}
