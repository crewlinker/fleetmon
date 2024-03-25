package maritime

import (
	"context"
)

const (
	// https://servicedocs.marinetraffic.com/tag/Search-Vessel
	pathVesselSearch = "shipsearch"
)

type VesselSearchService interface {
	VesselSearch(ctx context.Context, params VesselSearchParams) ([]VesselSearchResponse, error)
}

type VesselSearchResponse struct {
	serviceError ServiceError `json:"errors"`
	ShipName     string       `json:"SHIPNAME"`
	MMSI         string       `json:"MMSI"`
	IMO          string       `json:"IMO"`
	ShipID       string       `json:"SHIP_ID"`
	CallSign     string       `json:"CALLSIGN"`
	TypeName     string       `json:"TYPE_NAME"`
	DWT          string       `json:"DWT"`
	Flag         string       `json:"FLAG"`
	Country      string       `json:"COUNTRY"`
	YearBuilt    string       `json:"YEAR_BUILT"`
	MTURL        string       `json:"MT_URL"`
}

// VesselSearchParams represents the parameters for the VesselSearch method.
//
// One of ShipName, MMSI, IMO or ShipID is required.
type VesselSearchParams struct {
	ShipName   string `url:"shipname,omitempty"`     // Ship name (one of ShipName or MMSI or IMO or ShipID is required)
	MMSI       int    `url:"mmsi,omitempty"`         // MMSI Maritime Mobile Service Identity (one of ShipName or MMSI or IMO or ShipID is required)
	IMO        int    `url:"imo,omitempty"`          // IMO International Maritime Organization (one of ShipName or MMSI or IMO or ShipID is required)
	ShipID     int    `url:"shipid,omitempty"`       // ShipID (one of MMSI or IMO or ShipID is required)
	ShipType   int    `url:"shiptype,omitempty"`     // Ship type (optional)
	TypeNameID int    `url:"type_name_id,omitempty"` // Ship type name ID (optional)
	Protocol   string `url:"protocol,omitempty"`     // Protocol to use (xml, csv, json, jsono) (default: json)
}

func (v *VesselSearchParams) validate() error {
	if v.ShipName == "" && v.MMSI == 0 && v.IMO == 0 && v.ShipID == 0 {
		return ErrShipNameOrMMSIOrShipIDOrIMORequired
	}
	return nil
}

func (v *VesselSearchParams) setDefaults() {
	if v.Protocol == "" {
		v.Protocol = "json"
	}
}

func (c *Client) VesselSearch(ctx context.Context, params VesselSearchParams) ([]VesselSearchResponse, error) {
	if err := params.validate(); err != nil {
		return nil, err
	}

	params.setDefaults()

	req, err := c.newRequest(ctx, pathVesselSearch, params)
	if err != nil {
		return nil, err
	}

	var response []VesselSearchResponse
	if _, err := c.do(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
