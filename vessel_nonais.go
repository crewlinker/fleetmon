package fleetmon

import (
	"fmt"
)

// See: https://developer.fleetmon.com/reference/#operation/get-nonaisparticulars-detail
type VesselNonAISParameters struct {
	// Default: false
	// If true some extra attributes will be added to the response informing about usage and limits of this endpoint.
	RequestLimitInfo bool `url:"request_limit_info,omitempty"`
}

type VesselNonAISResponse struct {
	Error
	RequestLimitInfo RequestLimitInfo `json:"request_limit_info"`

	DeadWeight   int    `json:"dead_weight"`
	FleetmonURL  string `json:"fleetmon_url"`
	GrossTonnage int    `json:"gross_tonnage"`
	ImageURL     string `json:"image_url"`
	Manager      string `json:"manager"`
	Name         string `json:"name"`
	Owner        string `json:"owner"`
	VesselID     int    `json:"vessel_id"`
}

func (s *VesselService) NonAIS(vesselID int64, options VesselNonAISParameters) (*VesselNonAISResponse, error) {
	u, err := withOptions(fmt.Sprintf("vessel_nonais/%d", vesselID), options)
	if err != nil {
		return nil, err
	}

	req, err := s.client.NewRequest(u)
	if err != nil {
		return nil, err
	}

	r := new(VesselNonAISResponse)
	if resp, err := s.client.Do(req, r); err != nil {
		return nil, err
	} else if resp != nil && resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: %+v", r.Error)
	}

	return r, nil
}
