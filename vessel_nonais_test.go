// +build real

package fleetmon

import (
	"testing"
)

func TestVesselNonAIS(t *testing.T) {
	c := NewTestClient(t)
	ret, err := c.Vessel.NonAIS(46435, VesselNonAISParameters{})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("%+v", ret)

	if ret.DeadWeight < 100 || len(ret.FleetmonURL) == 0 {
		t.Fatalf("Something is wrong with the response")
	}
}
