package speSpectrum

import (
	avaatech "avaatechSpe"
)

type Spectrum struct {
	SPE    *avaatech.SPE    `json:"SPE"`
	Config *Configuration   `json:"-"`
	Peaks  []*Peak          `json:"-"`
	Lines  map[string]*Peak `json:"Lines"`
	Gain   float64          `json:"Gain"`
	Offset float64          `json:"Offset"`
	R2     float64          `json:"R2"`
}

type Peak struct {
	Channel float64 `json:"Channel"`
	Area    float64 `json:"Area"`
	FWHM    float64 `json:"FWHM"`
	ChiSq   float64 `json:"ChiSq"`
}

type Configuration struct {
	Threshold  float64 `json:"-"`
	GainMinKeV float64 `json:"-"`
	GainMaxKeV float64 `json:"-"`
}
