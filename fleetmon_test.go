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
