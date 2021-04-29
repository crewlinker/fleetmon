package fleetmon

import (
	"testing"
)

func TestNewRequest(t *testing.T) {
	u, err := withOptions("xyz", VesselSearchParameters{Callsign: "callsign"})
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	r, err := NewClient("abc").NewRequest(u)
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if r.URL.String() != "https://apiv2.fleetmon.com/xyz?callsign=callsign" {
		t.Fatalf("Invalid r.URL value: %s\n", r.URL)
	}
}

func TestFake(t *testing.T) {
	client := NewClient("testing")

	t.Run("basic", func(t *testing.T) {
		ret, err := client.Vessel.Basic(123, VesselBasicParameters{})
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		if ret == nil || ret.AisTypeOfShip == 0 {
			t.Fatalf("ooups: %+v", ret)
		}
	})

	t.Run("nonais", func(t *testing.T) {
		ret, err := client.Vessel.NonAIS(123, VesselNonAISParameters{})
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		if ret == nil || ret.DeadWeight == 0 {
			t.Fatalf("ooups: %+v", ret)
		}
	})

	t.Run("search", func(t *testing.T) {
		ret, err := client.Vessel.Search(VesselSearchParameters{})
		if err != nil {
			t.Fatalf("err: %s", err)
		}
		if ret == nil || len(ret.Vessels) == 0 {
			t.Fatalf("ooups: %+v", ret)
		}
	})
}
