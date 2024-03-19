package maritime

import "context"

const (
	// https://servicedocs.marinetraffic.com/tag/Vessel-Information#operation/exportvesselphoto
	pathVesselPhoto = "exportvesselphoto"
)

type VesselPhotoService interface {
	VesselPhoto(ctx context.Context, params VesselPhotoParams) ([]VesselPhotoResponse, error)
}

type VesselPhotoResponse struct {
	Photo string `json:"photo"`
}

type VesselPhotoParams struct {
	VesselID int    `url:"vessel_id,omitempty"` // IMO or MMSI number (one of MMSI or IMO or ShipID is required)
	ShipID   int    `url:"ship_id,omitempty"`   // Ship ID (one of MMSI or IMO or ShipID is required)
	Protocol string `url:"protocol,omitempty"`  // Protocol to use (json or xml) (default: json)
}

func (v *VesselPhotoParams) validate() error {
	if v.VesselID == 0 && v.ShipID == 0 {
		return ErrVesselIDOrShipIDRequired
	}
	return nil
}

func (v *VesselPhotoParams) setDefaults() {
	if v.Protocol == "" {
		v.Protocol = "json"
	}
}

func (c *Client) VesselPhoto(ctx context.Context, params VesselPhotoParams) ([]VesselPhotoResponse, error) {
	if err := params.validate(); err != nil {
		return nil, err
	}

	params.setDefaults()

	req, err := c.newRequest(ctx, pathVesselPhoto, params)
	if err != nil {
		return nil, err
	}

	var response []VesselPhotoResponse
	if _, err := c.do(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
