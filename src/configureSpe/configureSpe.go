package configureSpe

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Configuration struct {
	UTCoffset  string  `json:"-"`
	StdsPath   string  `json:"-"`
	Threshold  float64 `json:"-"`
	GainMinKeV float64 `json:"-"`
	GainMaxKeV float64 `json:"-"`
}

func ReadConfig() *Configuration {
	var fileRows, rowPts []string
	var key, value string

	config := new(Configuration)
	fileBytes, err := ioutil.ReadFile("configuration.cfg")
	if err != nil {
		log.Fatal(err)
	}
	fileStr := string(fileBytes)
	fileRows = strings.Split(fileStr, "\n")
	nRows := len(fileRows)
	for i := 0; i < nRows; i++ {
		rowPts = strings.Split(fileRows[i], "=")
		key = strings.Trim(rowPts[0], " ")
		value = strings.Trim(rowPts[1], " ")
		if key == "UTCoffset" {
			config.UTCoffset = value
		} else if key == "StdsPath" {
			config.StdsPath = value
		} else if key == "Threshold" {
			config.Threshold, err = strconv.ParseFloat(value, 64)
			if err != nil {
				config.Threshold = 1000.0
			}
		} else if key == "GainMinKeV" {
			config.GainMinKeV, err = strconv.ParseFloat(value, 64)
			if err != nil {
				config.GainMinKeV = 0.01980
			}
		} else if key == "GainMaxKeV" {
			config.GainMaxKeV, err = strconv.ParseFloat(value, 64)
			if err != nil {
				config.GainMaxKeV = 0.02050
			}
		}
	}
	return config
}
