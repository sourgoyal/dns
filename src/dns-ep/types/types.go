package types

// JSON structure mapping to be received from Drones (HTTP Client)
type DnsRequest struct {
	Xcord string `json:"x"`
	Ycord string `json:"y"`
	Zcord string `json:"z"`
	Vel   string `json:"vel"`
}

// JSON structure mapping to be sent to Drones (HTTP client)
type DnsLocResp struct {
	Location float64 `json:"loc"`
}
