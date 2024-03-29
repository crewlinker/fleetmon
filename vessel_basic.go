package fleetmon

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

// See: https://developer.fleetmon.com/reference/#operation/get-vesselbasic-detail
type VesselBasicParameters struct {
	// Default: false
	// If true some extra attributes will be added to the response informing about usage and limits of this endpoint.
	RequestLimitInfo bool `url:"request_limit_info,omitempty"`
}

type VesselBasicResponse struct {
	Error
	RequestLimitInfo RequestLimitInfo `json:"request_limit_info"`

	AisTypeOfShip    int    `json:"ais_type_of_ship"`
	AisTypeOfShipStr string `json:"ais_type_of_ship_str"`
	Callsign         string `json:"callsign"`
	CnIso2           string `json:"cn_iso2"`
	CnName           string `json:"cn_name"`
	FleetmonURL      string `json:"fleetmon_url"`
	ImageURL         string `json:"image_url"`
	ImoNumber        int    `json:"imo_number"`
	Length           int    `json:"length"`
	MmsiNumber       int    `json:"mmsi_number"`
	Name             string `json:"name"`
	Type             string `json:"type"`
	TypeClass        string `json:"type_class"`
	TypeCode         string `json:"type_code"`
	VesselID         int    `json:"vessel_id"`
	Width            int    `json:"width"`
}

//go:embed fixtures/basic_46435.json
var basicFixture []byte

// XXX this is not tested against the real api
func (s *VesselService) Basic(vesselID int64, options VesselBasicParameters) (*VesselBasicResponse, error) {
	r := new(VesselBasicResponse)
	if s.client.isTesting() {
		if err := json.Unmarshal(basicFixture, r); err != nil {
			return nil, err
		}
		return r, nil
	}

	u, err := withOptions(fmt.Sprintf("basicvessel/%d", vesselID), options)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(u)
	if err != nil {
		return nil, err
	}

	if resp, err := s.client.Do(req, r); err != nil {
		return nil, err
	} else if resp != nil && resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: %+v", r.Error)
	}

	return r, nil
}
