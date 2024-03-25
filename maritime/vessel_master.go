package maritime

import (
	"context"
)

const (
	vesselMasterServiceVersion = 5 // Default version to be used in the VesselMasterService
	pathVesselMaster           = "vesselmasterdata"
)

type VesselMasterService interface {
	VesselMaster(ctx context.Context, params VesselMasterParameters) ([]VesselMasterResponse, error)
}

// VesselMasterResponse represents the response from the VesselMasterService
type VesselMasterResponse struct {
	MMSI                   string      `json:"MMSI"`
	IMO                    string      `json:"IMO"`
	ENI                    interface{} `json:"ENI"`
	Name                   string      `json:"NAME"`
	Builder                string      `json:"BUILDER"`
	PlaceOfBuild           string      `json:"PLACE_OF_BUILD"`
	Build                  string      `json:"BUILD"`
	YardNumber             string      `json:"YARD_NUMBER"`
	BreadthExtreme         string      `json:"BREADTH_EXTREME"`
	BreadthMoulded         string      `json:"BREADTH_MOULDED"`
	Depth                  string      `json:"DEPTH"`
	SummerDWT              string      `json:"SUMMER_DWT"`
	DisplacementSummer     interface{} `json:"DISPLACEMENT_SUMMER"`
	CallSign               string      `json:"CALLSIGN"`
	Flag                   string      `json:"FLAG"`
	Draught                string      `json:"DRAUGHT"`
	LengthOverall          string      `json:"LENGTH_OVERALL"`
	LengthBWPerpendiculars string      `json:"LENGTH_B_W_PERPENDICULARS"`
	FuelConsumption        string      `json:"FUEL_CONSUMPTION"`
	SpeedMax               interface{} `json:"SPEED_MAX"`
	SpeedService           string      `json:"SPEED_SERVICE"`
	TEU                    string      `json:"TEU"`
	GrossTonnage           string      `json:"GROSS_TONNAGE"`
	NetTonnage             string      `json:"NET_TONNAGE"`
	LiquidOil              interface{} `json:"LIQUID_OIL"`
	LiquidGas              interface{} `json:"LIQUID_GAS"`
	Grain                  interface{} `json:"GRAIN"`
	Owner                  string      `json:"OWNER"`
	Manager                string      `json:"MANAGER"`
	FinancialOwner         string      `json:"FINANCIAL_OWNER"`
	TechnicalManager       interface{} `json:"TECHNICAL_MANAGER"`
	Insurer                string      `json:"INSURER"`
	VesselType             string      `json:"VESSEL_TYPE"`
}

// VesselMasterParameters represents the parameters to be used in the VesselMasterService.
//
// One of MMSI, IMO or ShipID is required.
type VesselMasterParameters struct {
	Version     int    `url:"v"`                  // Version of the API to use (default: 5)
	MMSI        int    `url:"mmsi,omitempty"`     // MMSI Maritime Mobile Service Identity (one of MMSI or IMO or ShipID is required)
	IMO         int    `url:"imo,omitempty"`      // IMO International Maritime Organization (one of MMSI or IMO or ShipID is required)
	ShipID      int    `url:"shipid,omitempty"`   // ShipID (one of MMSI or IMO or ShipID is required)
	Timespan    int    `url:"timespan,omitempty"` // Timespan of last received position signal (minutes: max 2880, days: 5)  (optional)
	Interval    string `url:"interval,omitempty"` // Interval of last received position signal (minutes, days) (default: days)
	Protocol    string `url:"protocol,omitempty"` // Protocol to use (xml, csv, json, jsono) (default: json)
	Page        int    `url:"page,omitempty"`     // Page number (default: 1)
	MessageType string `url:"msgtype,omitempty"`  // MessageType (simple, extended) (default: simple)
}

// validate checks if the VesselMasterParameters are valid
func (v *VesselMasterParameters) validate() error {
	if v.MMSI == 0 && v.IMO == 0 && v.ShipID == 0 {
		return ErrVesselIDOrShipIDRequired
	}

	return nil
}

// setDefaults sets the default values for the VesselMasterParameters
func (v *VesselMasterParameters) setDefaults() {
	if v.Version == 0 {
		v.Version = vesselMasterServiceVersion
	}
	if v.Protocol == "" {
		v.Protocol = "json"
	}
}

// VesselMaster returns the master data of a vessel
func (c *Client) VesselMaster(ctx context.Context, params VesselMasterParameters) ([]VesselMasterResponse, error) {
	// validate the parameters
	if err := params.validate(); err != nil {
		return nil, err
	}

	// set the default values
	params.setDefaults()

	req, err := c.newRequest(ctx, pathVesselMaster, params)
	if err != nil {
		return nil, err
	}

	var resp []VesselMasterResponse
	if _, err := c.do(req, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}
