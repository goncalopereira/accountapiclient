package e2e

import (
	. "interview-accountapi/internal/http"
	"testing"
)

//	"github.com/stretchr/testify/assert"

func TestHealth(t *testing.T) {
	err := Health()

	if err != nil {
		t.Errorf("Healthcheck failed with error: %v", err)
	}
}
