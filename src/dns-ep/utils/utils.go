package utils

import (
	"log"
	"strconv"
)

func StrConvFloat(strX, strY, strZ, strVel string) (x, y, z, vel float64, err error) {

	x, err = strconv.ParseFloat(strX, 32)
	if err != nil {
		log.Print("ParseFloat failed", err)
		return 0, 0, 0, 0, err
	}
	y, err = strconv.ParseFloat(strY, 32)
	if err != nil {
		log.Print("ParseFloat failed", err)
		return 0, 0, 0, 0, err
	}
	z, err = strconv.ParseFloat(strZ, 32)
	if err != nil {
		log.Print("ParseFloat failed", err)
		return 0, 0, 0, 0, err
	}
	vel, err = strconv.ParseFloat(strVel, 32)
	if err != nil {
		log.Print("ParseFloat failed", err)
		return 0, 0, 0, 0, err
	}

	return x, y, z, vel, nil
}
