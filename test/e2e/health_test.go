package e2e

import (
	http "github.com/goncalopereira/accountapiclient/internal/commands"
	"testing"
)

//	"github.com/stretchr/testify/assert"

func TestHealth(t *testing.T) {
	err := http.Health()

	if err != nil {
		t.Errorf("Healthcheck failed with error: %v", err)
	}
}
