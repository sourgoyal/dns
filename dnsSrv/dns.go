package dns

// DNS Info data structure
type DnsInfo struct {
	SectorId int
	Info     DnsOper
}

// DNS Operations
type DnsOper interface {
	GetSectorId() int
	SetSectorId(int) error
	CalcLocation(DnsReq) (float64, error)
}

// DNS request params
type DnsReq struct {
	X   float64
	Y   float64
	Z   float64
	Vel float64
}

// Create a new instance of DNS
func New() *DnsInfo {
	return new(DnsInfo)
}

// API to get Sector ID
func (dInfo *DnsInfo) GetSectorId() int {
	return dInfo.SectorId
}

// API to set Sector ID
func (dInfo *DnsInfo) SetSectorId(id int) error {
	dInfo.SectorId = id
	return nil
}

// Bussiness logic API, to calculate location based on sectorID, cordinates and Velocity.
func (dInfo *DnsInfo) CalcLocation(req *DnsReq) (float64, error) {
	sectorId := float64(dInfo.SectorId)
	return sectorId*req.X + sectorId*req.Y + sectorId*req.Z + req.Vel, nil
}
