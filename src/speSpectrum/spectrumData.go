package speSpectrum

import (
	avaatech "avaatechSpe"
)

type Spectrum struct {
	SPE   *avaatech.SPE
	Peaks []*Peak
}

type Peak struct {
	Channel int64
	Area    int64
	FWHM    int64
}
