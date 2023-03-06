package runners

import (
	"testing"
)

func TestRunners(t *testing.T) {
	result := runners("works")
	if result != "runners works" {
		t.Error("Expected runners to append 'works'")
	}
}
