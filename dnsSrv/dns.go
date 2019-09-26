package dns

type DnsInfo struct {
	SectorId int
	Info     DnsOper
}

type DnsOper interface {
	GetSectorId() int
	SetSectorId(int) error
	CalcLocation(DnsReq) (float64, error)
}

type DnsReq struct {
	X   float64
	Y   float64
	Z   float64
	Vel float64
}

func New() *DnsInfo {
	return new(DnsInfo)
}

func (dInfo *DnsInfo) GetSectorId() int {
	return dInfo.SectorId
}

func (dInfo *DnsInfo) SetSectorId(id int) error {
	dInfo.SectorId = id
	return nil
}

func (dInfo *DnsInfo) CalcLocation(req *DnsReq) (float64, error) {
	sectorId := float64(dInfo.SectorId)
	return sectorId*req.X + sectorId*req.Y + sectorId*req.Z + req.Vel, nil
}
