package speSpectrum

import (
	avaatech "avaatechSpe"
)

type Spectrum struct {
	SPE     *avaatech.SPE
	PeakMin int64
	Peaks   [][]int64
}
