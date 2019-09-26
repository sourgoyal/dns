package main

import (
	"bytes"
	"dns-ep/types"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestDns(t *testing.T) {
	var wg sync.WaitGroup

	RequestBurstChan := make(chan struct{}, 1)

	for i := 0; i < 10; i++ {
		RequestBurstChan <- struct{}{}
		wg.Add(1)
		go func() {
			defer func() { time.Sleep(100 * time.Millisecond); <-RequestBurstChan }()
			defer wg.Done()
			url := "http://localhost:8080/getLoc"
			fmt.Println("URL:>", url)

			jsonStr, err := json.Marshal(types.DnsRequest{Xcord: "123.12", Ycord: "456.56", Zcord: "789.89", Vel: "20"})
			if err != nil {
				t.Error("Marshing failed", err)
			}
			log.Printf("jsonStr %s\n", jsonStr)

			// var jsonStr = []byte(`{"x": "22.11", "y": "22.11", "z": "22.11", "vel": "22.11"}`)

			req, er := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
			if er != nil {
				t.Error("http.NewRequest failed", er)
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				t.Error("http failed", err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Error("ReadAll failed", err)
			}
			fmt.Println("Response: ", string(body))
		}()
	}
	wg.Wait()
}
