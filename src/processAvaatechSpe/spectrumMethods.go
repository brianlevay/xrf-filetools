package processAvaatechSpe

import ()

// Right now, this only gets the maximum intensities of the peaks and the height relative to
// the nearest saddle points. A previous attempt to model the peaks with Gaussian curves
// produced very noisy results and was dropped. I might try again another time.

func (spect *Spectrum) ModelPeaks(threshold float64) {
	inflections := getInflections(spect.SPE.Counts)
	peakList := getPeaks(inflections, threshold)
	spect.Peaks = peakList
}

// Tests with bAxil on spectra from both instruments indicate that offset is
// always << 20 eV, which is < 1 channel width. Thus, it can be
// approximated as 0 eV for all fitting in this module.
// Gain varies from ~20.0 eV/ch to ~20.4 eV/ch, which leads to channel differences of
// ~2 at Si Ka and ~7 out at Fe Kb. Thus, gain needs to be a variable. Min and Max
// allowable gain values will be passed through when calling "Process"

func (spect *Spectrum) AssignPeaks(gainMinKeV float64, gainMaxKeV float64) {
	var lineMap map[string]*Peak
	var peak *Peak
	var startSearch int
	var maxLineCh, minLineCh float64
	lineMap = make(map[string]*Peak)
	nPeaks := len(spect.Peaks)
	nLines := len(lineList)
	startSearch = 0
	for i := 0; i < nPeaks; i++ {
		peak = spect.Peaks[i]
		for j := startSearch; j < nLines; j++ {
			minLineCh = lineList[j].Energy / gainMaxKeV // maxGain => minCh
			maxLineCh = lineList[j].Energy / gainMinKeV
			if (peak.Channel >= minLineCh) && (peak.Channel <= maxLineCh) {
				lineMap[lineList[j].Name] = peak
				startSearch = j + 1
				break
			}
		}
	}
	spect.Lines = lineMap
}
