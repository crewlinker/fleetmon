package fleetmon

import (
	"fmt"
	"time"
)

// See: https://developer.fleetmon.com/reference/#operation/get-vesselsearch-listing
type VesselSearchParameters struct {
	// only exact matching
	Callsign string `url:"callsign,omitempty"`

	// only exact matching
	ENINumber int64 `url:"eni_number,omitempty"`

	// only exact matching
	IMONumber int64 `url:"imo_number,omitempty"`

	// maximum timespan of last received position signal, in days
	MaxPosAgeDays int64 `url:"max_pos_age_days,omitempty"`

	// only exact matching
	MMSINumber int64 `url:"mmsi_number,omitempty"`

	// sub-string matching, minium length is 3 characters
	Name string `url:"name,omitempty"`

	// Default: false
	// If true some extra attributes will be added to the response informing about usage and limits of this endpoint.
	RequestLimitInfo bool `url:"request_limit_info,omitempty"`

	// Default: true
	// request only vessels with a valid position
	WithPosition bool `url:"with_position,omitempty"`
}

type Vessel struct {
	Callsign         string    `json:"callsign"`
	CnIso2           string    `json:"cn_iso2"`
	CnName           string    `json:"cn_name"`
	ImoNumber        int       `json:"imo_number"`
	Length           int       `json:"length"`
	MmsiNumber       int       `json:"mmsi_number"`
	Name             string    `json:"name"`
	PositionReceived time.Time `json:"position_received"`
	StaticReceived   time.Time `json:"static_received"`
	VesselID         int       `json:"vessel_id"`
	VtVerbose        string    `json:"vt_verbose"`
}

type VesselSearchResponse struct {
	Error
	RequestLimitInfo RequestLimitInfo `json:"request_limit_info"`

	Vessels []Vessel `json:"vessels"`
}

func (s *VesselService) Search(options VesselSearchParameters) (*VesselSearchResponse, error) {
	u, err := withOptions("vesselsearch", options)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(u)
	if err != nil {
		return nil, err
	}

	r := new(VesselSearchResponse)
	if resp, err := s.client.Do(req, r); err != nil {
		return nil, err
	} else if resp != nil && resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: %+v", r.Error)
	}

	return r, nil
}
