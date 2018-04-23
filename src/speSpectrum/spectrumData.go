package speSpectrum

import (
	avaatech "avaatechSpe"
)

type Spectrum struct {
	SPE    *avaatech.SPE
	Config *Configuration
	Peaks  []*Peak
	Lines  map[string]*Peak
	Gain   float64
	Offset float64
	R2     float64
}

type Peak struct {
	Channel float64
	Area    float64
	FWHM    float64
	ChiSq   float64
}

type Configuration struct {
	Threshold  float64
	GainMinKeV float64
	GainMaxKeV float64
}
