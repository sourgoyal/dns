package dns

import "testing"

func TestSetGetSectorId(t *testing.T) {
	var dns *DnsInfo
	if dns = New(); dns == nil {
		t.Error("DNS init failed")
		return
	}

	if err := dns.SetSectorId(10); err != nil {
		t.Error("Failed to set SectorID")
	}

	if dns.GetSectorId() != 10 {
		t.Errorf("Invalid Sector ID %d", dns.GetSectorId())
	}
}
