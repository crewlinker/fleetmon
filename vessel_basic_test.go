// +build real

package fleetmon

import (
	"testing"
)

func TestVesselBasic(t *testing.T) {
	c := NewTestClient(t)
	ret, err := c.Vessel.Basic(46435, VesselBasicParameters{})
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	t.Logf("%+v", ret)

	if ret.Length < 10 || len(ret.Callsign) == 0 {
		t.Fatalf("Something is wrong with the response")
	}
}
