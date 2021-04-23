package fleetmon

type RequestLimitInfo struct {
	LeftRequests int `json:"left_requests"`
	MaxRequests  int `json:"max_requests"`
	UsedRequests int `json:"used_requests"`
}

type Error struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
	Message     string `json:"message"`
	Status      string `json:"status"`
}
