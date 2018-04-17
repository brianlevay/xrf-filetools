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
	Height  int64
	Width   int64
}
