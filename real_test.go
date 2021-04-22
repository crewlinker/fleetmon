// +build real

package fleetmon

import (
	"os"
	"testing"
)

func NewTestClient(t *testing.T) *Client {
	t.Helper()

	key := os.Getenv("FLEETMON_KEY")

	if len(key) == 0 {
		t.Fatalf("FLEETMON_KEY unset")
	}

	return NewClient(key)
}
