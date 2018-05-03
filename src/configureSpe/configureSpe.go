package configureSpe

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Configuration struct {
	UTCoffset     string  `json:"-"`
	MaxChannel    int     `json:"-"`
	SNIPwidth     float64 `json:"-"`
	MinPeakHeight float64 `json:"-"`
	GainMinKeV    float64 `json:"-"`
	GainMidKeV    float64 `json:"-"`
	GainMaxKeV    float64 `json:"-"`
}

func (config *Configuration) SetDefaults() {
	config.UTCoffset = "0"
	config.MaxChannel = 500
	config.SNIPwidth = 50.0
	config.MinPeakHeight = 500.0
	config.GainMinKeV = 0.01980
	config.GainMidKeV = 0.02000
	config.GainMaxKeV = 0.02050
}

func ReadConfig() (string, *Configuration) {
	var fileRows, rowPts []string
	var key, value string
	var stdsPath string
	var maxChannel int
	var snipWidth, minPeakHeight, gainMinKeV, gainMidKeV, gainMaxKeV float64

	config := new(Configuration)
	config.SetDefaults()
	fileBytes, err := ioutil.ReadFile("configuration.cfg")
	if err != nil {
		log.Fatal(err)
	}
	fileStr := string(fileBytes)
	fileRows = strings.Split(fileStr, "\n")
	nRows := len(fileRows)
	for i := 0; i < nRows; i++ {
		if strings.Contains(fileRows[i], "!") == false {
			rowPts = strings.Split(fileRows[i], "=")
			key = strings.Trim(rowPts[0], " ")
			value = strings.Trim(rowPts[1], " ")
			if key == "StdsPath" {
				stdsPath = value
			} else if key == "UTCoffset" {
				config.UTCoffset = value
			} else if key == "MaxChannel" {
				maxChannel, err = strconv.Atoi(value)
				if err == nil {
					config.MaxChannel = maxChannel
				}
			} else if key == "SNIPwidth" {
				snipWidth, err = strconv.ParseFloat(value, 64)
				if err == nil {
					config.SNIPwidth = snipWidth
				}
			} else if key == "MinPeakHeight" {
				minPeakHeight, err = strconv.ParseFloat(value, 64)
				if err == nil {
					config.MinPeakHeight = minPeakHeight
				}
			} else if key == "GainMinKeV" {
				gainMinKeV, err = strconv.ParseFloat(value, 64)
				if err == nil {
					config.GainMinKeV = gainMinKeV
				}
			} else if key == "GainMidKeV" {
				gainMidKeV, err = strconv.ParseFloat(value, 64)
				if err == nil {
					config.GainMidKeV = gainMidKeV
				}
			} else if key == "GainMaxKeV" {
				gainMaxKeV, err = strconv.ParseFloat(value, 64)
				if err != nil {
					config.GainMaxKeV = gainMaxKeV
				}
			}
		}
	}
	return stdsPath, config
}
