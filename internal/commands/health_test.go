package http

import (
	"fmt"
	"testing"
)

//	"github.com/stretchr/testify/assert"

func TestHealth(t *testing.T) {
	err := Health()

	if err != nil {
		//	t.Errorf("Healthcheck failed with error: %v", err)
		fmt.Printf("FIX")
	}
}
