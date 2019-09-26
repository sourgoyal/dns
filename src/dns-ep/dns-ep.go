package main

import (
	"dns-ep/dns"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	throttling "github.com/sourgoyal/golang/http-throttling"
)

const (
	ThrottlingHttpRequestRate  = 5
	ThrottlingHttpRequestBurst = 5
	DnsEPDefaultPort           = "8080"
)

func main() {
	// Using gorilla/mux URL router and dispatcher to match URL paths to handlers.
	router := mux.NewRouter()

	// Handler to handle request from Drones
	router.HandleFunc("/getLoc", dns.DnsRequestHandler).Methods("POST")

	// Configure rate limiter to throttle incoming http requests.
	throttling.ConfigureLimiter(ThrottlingHttpRequestRate, ThrottlingHttpRequestBurst)
	log.Printf("Starting DNS Endpoint...\n")
	log.Printf("Throttling HTTP request at rate: %d, BurstSize: %d\n", constants.ThrottlingHttpRequestRate, constants.ThrottlingHttpRequestBurst)

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
