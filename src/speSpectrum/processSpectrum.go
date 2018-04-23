package speSpectrum

import (
	avaatech "avaatechSpe"
)

func Process(spe *avaatech.SPE) *Spectrum {
	spect := new(Spectrum)
	spect.SPE = spe
	spect.Threshold = peakCutoff
	spect.ModelPeaks()
	spect.AssignPeaks()
	spect.CalculateEnergyScale()
	return spect
}

func (spect *Spectrum) ModelPeaks() {
	var peakList []*Peak
	var peak *Peak
	inflections := getInflections(spect.SPE.Counts)
	peakPositions := getPeakPositions(inflections, spect.Threshold)
	nPeaks := len(peakPositions)
	for i := 0; i < nPeaks; i++ {
		peak = fitPeakLinearSearch(peakPositions[i], spect.SPE.Counts)
		peakList = append(peakList, peak)
	}
	spect.Peaks = peakList
}

func (spect *Spectrum) AssignPeaks() {
	var peak *Peak
	var lastAssigned int
	var maxGain, minGain, maxLineCh, minLineCh float64
	spect.Lines = make(map[string]*Peak)
	maxGain = gain_keV + gain_delta
	minGain = gain_keV - gain_delta
	nPeaks := len(spect.Peaks)
	nLines := len(line_keV)
	lastAssigned = 0
	for i := 0; i < nPeaks; i++ {
		peak = spect.Peaks[i]
		for j := lastAssigned; j < nLines; j++ {
			maxLineCh = channelFromEnergy(line_keV[j].Energy, maxGain, offset_keV)
			minLineCh = channelFromEnergy(line_keV[j].Energy, minGain, offset_keV)
			if (peak.Channel >= minLineCh) || (peak.Channel <= maxLineCh) {
				spect.Lines[line_keV[j].Name] = peak
				lastAssigned = j
				break
			}
		}
	}
}

func (spect *Spectrum) CalculateEnergyScale() {

}
