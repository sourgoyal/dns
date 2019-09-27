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
	SectorId                   = 1   // SectorId is taken constant as requirement
	ThrottlingHttpRequestRate  = 100 // Rate limit HTTP requests per second. Client shall retry if failed to get location
	ThrottlingHttpRequestBurst = 50  // Burst size allowed
	DnsEPDefaultPort           = "8080"
)

// DNS holds bussiness logic for this app.
// It can be used with any REST, gRPC, etc
// It can be modified into map[sectorID] with RW Lock in case, this DNS has to serve multiple sectors
var Dns *dns.DnsInfo

func main() {
	// Configure DNS and assign sectorID to it.
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

// Handler for "/getLoc"
func dnsRequestHandler(w http.ResponseWriter, r *http.Request) {
	dnsReq := types.DnsRequest{}

	// Read JSON from the HTTP request body
	jsn, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println("HttpRequestReadFailed " + err.Error())
		return
	}

	// Unmarshal JSON into DnsRequest{} structure
	err = json.Unmarshal(jsn, &dnsReq)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println("UnmarshallingFailed " + err.Error())
		return
	}
	log.Printf("JSON received %s\n", jsn)

	// Convert string into floats
	xCordFloat, yCordFloat, zCordFloat, velFloat, err := utils.StrConvFloat(dnsReq.Xcord, dnsReq.Ycord, dnsReq.Zcord, dnsReq.Vel)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println("ParseStringToFloatFailed " + err.Error())
		return
	}

	// Call DNS to get the locations for the HTTP request
	dnsR := &dns.DnsReq{X: xCordFloat, Y: yCordFloat, Z: zCordFloat, Vel: velFloat}
	location, err := Dns.CalcLocation(dnsR)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println("Location Calculation Failed " + err.Error())
		return
	}

	// Prepare JSON response to be sent to HTTP client as response
	loc := types.DnsLocResp{Location: location}
	locJsn, er := json.Marshal(loc)
	if er != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println("MarshallingRespFailed " + er.Error())
		return
	}

	utils.SendJsonToClient(w, locJsn)
}
