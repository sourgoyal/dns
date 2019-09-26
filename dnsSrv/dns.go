package dnsSrv

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"../dnsSrv/constants"
)

type DnsRequest struct {
	Xcord string `json:"x"`
	Ycord string `json:"y"`
	Zcord string `json:"z"`
	Vel   string `json:"vel"`
}

type DnsLocResp struct {
	Location float32 `json:"loc"`
}

func DnsRequestHandler(w http.ResponseWriter, r *http.Request) {
	dnsReq := DnsRequest{}

	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading the request body", err)
	}

	err = json.Unmarshal(jsn, &dnsReq)
	if err != nil {
		log.Fatal("Unmarshal failed", err)
	}

	log.Printf("Body received %s\n", jsn)
	log.Printf("dnsReq %+v\n", dnsReq)

	xCordFloat, err := strconv.ParseFloat(dnsReq.Xcord, 32)
	if err != nil {
		log.Fatal("ParseFloat failed", err)
	}
	yCordFloat, err := strconv.ParseFloat(dnsReq.Ycord, 32)
	if err != nil {
		log.Fatal("ParseFloat failed", err)
	}
	zCordFloat, err := strconv.ParseFloat(dnsReq.Zcord, 32)
	if err != nil {
		log.Fatal("ParseFloat failed", err)
	}
	velFloat, err := strconv.ParseFloat(dnsReq.Vel, 32)
	if err != nil {
		log.Fatal("ParseFloat failed", err)
	}

	sectorId := float64(constants.SectorId)
	location := sectorId*xCordFloat + sectorId*yCordFloat + sectorId*zCordFloat + velFloat

	loc := DnsLocResp{Location: float32(location)}
	locJsn, er := json.Marshal(loc)
	if er != nil {
		log.Fatal(constants.MarshalingFailedErr, err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(locJsn)
}
