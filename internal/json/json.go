package json

import (
	"encoding/json"
	"fmt"
)

func DataToBytes(data interface{}) ([]byte, error) {
	bytes, err := json.Marshal(data)

	if err != nil {
		return nil, fmt.Errorf("DataToBytes: %w", err)
	}

	return bytes, nil
}
