package speSpectrum

import (
	avaatech "avaatechSpe"
	"fmt"
)

const bgMult = 3

func GetSpectrum(spe *avaatech.SPE) *Spectrum {
	spect := new(Spectrum)
	spect.MetaData(spe)
	peaks := getPeaks(spe.Counts, bgMult)
	fmt.Println(len(peaks)) //// FOR DEBUGGING ////
	return spect
}
