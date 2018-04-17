package speSpectrum

import (
	avaatech "avaatechSpe"
)

type Spectrum struct {
	SPE   *avaatech.SPE
	Peaks []*Peak
}

type Peak struct {
	Channel float64
	Area    float64
	FWHM    float64
	ChiSq   float64
}
