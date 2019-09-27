package utils

import (
	"log"
	"net/http"
	"strconv"
)

// Util API to convert string x, y, z and vel parameters into float64
func StrConvFloat(strX, strY, strZ, strVel string) (x, y, z, vel float64, err error) {

	x, err = strconv.ParseFloat(strX, 8)
	if err != nil {
		log.Print("ParseFloat failed", err)
		return 0, 0, 0, 0, err
	}
	y, err = strconv.ParseFloat(strY, 8)
	if err != nil {
		log.Print("ParseFloat failed", err)
		return 0, 0, 0, 0, err
	}
	z, err = strconv.ParseFloat(strZ, 8)
	if err != nil {
		log.Print("ParseFloat failed", err)
		return 0, 0, 0, 0, err
	}
	vel, err = strconv.ParseFloat(strVel, 8)
	if err != nil {
		log.Print("ParseFloat failed", err)
		return 0, 0, 0, 0, err
	}

	return x, y, z, vel, nil
}

func SendJsonToClient(w http.ResponseWriter, json []byte) {
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(json); err != nil {
		log.Printf("Error Occured while sending error to client err: %+v", err)
	}
}
