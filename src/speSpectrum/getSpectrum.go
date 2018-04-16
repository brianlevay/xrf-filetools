package speSpectrum

import (
	avaatech "avaatechSpe"
)

const peakMin int64 = 1000

func GetSpectrum(spe *avaatech.SPE) *Spectrum {
	spect := new(Spectrum)
	spect.SPE = spe
	spect.PeakMin = peakMin
	spect.GetPeaks()
	return spect
}
