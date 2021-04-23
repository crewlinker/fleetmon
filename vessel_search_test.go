// +build real

package fleetmon

import (
	"testing"
)

func TestVesselSearch(t *testing.T) {
	c := NewTestClient(t)
	ret, err := c.Vessel.Search(VesselSearchParameters{Name: "fairway"})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("%+v", ret)

	if len(ret.Vessels) < 8 {
		t.Fatalf("There should be at least 8 vessels matching this query, have %d", len(ret.Vessels))
	}
}
