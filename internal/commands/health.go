package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Health() error {
	response, responseErr := http.Get("http://localhost:8080/v1/health")

	if response != nil {
		defer response.Body.Close()
	}
	if responseErr != nil {
		return responseErr
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("accounts health HTTP status: %v", response.StatusCode)
	}

	body, bodyErr := ioutil.ReadAll(response.Body)

	if responseErr != nil {
		return fmt.Errorf("accounts health error body: %v", bodyErr.Error())
	}

	formattedBody := string(body)
	if formattedBody != "{\"status\":\"up\"}" {
		return fmt.Errorf("accounts health error body: %v", formattedBody)
	}
	return nil
}
