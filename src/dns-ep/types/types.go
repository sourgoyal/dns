package types

type DnsRequest struct {
	Xcord string `json:"x"`
	Ycord string `json:"y"`
	Zcord string `json:"z"`
	Vel   string `json:"vel"`
}

type DnsLocResp struct {
	Location float32 `json:"loc"`
}
