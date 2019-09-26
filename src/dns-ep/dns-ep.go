package main

import (
	"dns-ep/types"
	"dns-ep/utils"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	dns "github.com/sourgoyal/dns/dnsSrv"

	"github.com/gorilla/mux"
	throttling "github.com/sourgoyal/golang/http-throttling"
)

const (
	SectorId                   = 1
	ThrottlingHttpRequestRate  = 5
	ThrottlingHttpRequestBurst = 5
	DnsEPDefaultPort           = "8080"
)

var Dns *dns.DnsInfo

func main() {
	// Configure DNS
	Dns = dns.New()
	if err := Dns.SetSectorId(SectorId); err != nil {
		log.Fatal("Failed to initialize...")
	}

	// Using gorilla/mux URL router and dispatcher to match URL paths to handlers.
	router := mux.NewRouter()

	// Handler to handle request from Drones
	router.HandleFunc("/getLoc", dnsRequestHandler).Methods("POST")

	// Configure rate limiter to throttle incoming http requests.
	throttling.ConfigureLimiter(ThrottlingHttpRequestRate, ThrottlingHttpRequestBurst)
	log.Printf("Starting DNS Endpoint...\n")
	log.Printf("Throttling HTTP request at rate: %d, BurstSize: %d\n", ThrottlingHttpRequestRate, ThrottlingHttpRequestBurst)

	// Use exported DNSPORT
	port := os.Getenv("DNSPORT")
	if len(port) <= 0 {
		port = DnsEPDefaultPort
	}
	log.Printf("Using PORT: %s", port)

	if err := http.ListenAndServe(":"+port, throttling.LimitRate(router)); err != nil {
		log.Fatal("DNS EP failed to start", err)
	}
}

func dnsRequestHandler(w http.ResponseWriter, r *http.Request) {
	dnsReq := types.DnsRequest{}

	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		sendErr(w, "HttpRequestReadFailed "+err.Error())
		return
	}

	err = json.Unmarshal(jsn, &dnsReq)
	if err != nil {
		sendErr(w, "UnmarshallingFailed "+err.Error())
		return
	}

	log.Printf("Body received %s\n", jsn)
	log.Printf("dnsReq %+v\n", dnsReq)

	xCordFloat, yCordFloat, zCordFloat, velFloat, err := utils.StrConvFloat(dnsReq.Xcord, dnsReq.Ycord, dnsReq.Zcord, dnsReq.Vel)
	if err != nil {
		sendErr(w, "ParseStringToFloatFailed "+err.Error())
		return
	}

	dnsR := &dns.DnsReq{X: xCordFloat, Y: yCordFloat, Z: zCordFloat, Vel: velFloat}
	location, err := Dns.CalcLocation(dnsR)
	if err != nil {
		sendErr(w, "LocationFailed "+err.Error())
		return
	}

	loc := types.DnsLocResp{Location: float32(location)}
	locJsn, er := json.Marshal(loc)
	if err != nil {
		sendErr(w, "MarshallingRespFailed "+er.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err = w.Write(locJsn); err != nil {
		log.Printf("Error Occured while sending error to client err: %+v", err)
	}
}

func sendErr(w http.ResponseWriter, err string) {
	if _, e := w.Write([]byte(err)); e != nil {
		log.Print("Error Occured while sending error to client")
	}
}
