package speSpectrum

import (
	avaatech "avaatechSpe"
)

func Process(spe *avaatech.SPE) *Spectrum {
	spect := new(Spectrum)
	spect.SPE = spe
	spect.ModelPeaks()
	spect.AssignPeaks()
	return spect
}

func (spect *Spectrum) ModelPeaks() {
	var peakList []*Peak
	var peak *Peak
	inflections := getInflections(spect.SPE.Counts)
	peakPositions := getPeakPositions(inflections)
	nPeaks := len(peakPositions)
	for i := 0; i < nPeaks; i++ {
		peak = fitPeakLinearSearch(peakPositions[i], spect.SPE.Counts)
		peakList = append(peakList, peak)
	}
	spect.Peaks = peakList
}

func (spect *Spectrum) AssignPeaks() {

}
