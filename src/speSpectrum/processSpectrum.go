package speSpectrum

import (
	avaatech "avaatechSpe"
)

func Process(spe *avaatech.SPE) *Spectrum {
	spect := new(Spectrum)
	spect.SPE = spe
	spect.GetPeaks(peakCutoff)
	return spect
}
