package processAvaatechSpe

import (
	"math"
)

func (spect *Spectrum) ModelPeaks(threshold float64) {
	var peakList []*Peak
	var peak *Peak
	inflections := getInflections(spect.SPE.Counts)
	peakPositions := getPeakPositions(inflections, threshold)
	nPeaks := len(peakPositions)
	for i := 0; i < nPeaks; i++ {
		peak = fitPeakLinearSearch(peakPositions[i], spect.SPE.Counts)
		peakList = append(peakList, peak)
	}
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

func (spect *Spectrum) CalculateEnergyScale() {
	var sumX, sumY, sumXX, sumXY, sumYY float64
	var aveX, aveY float64
	var Sxx, Syy, Sxy float64
	var slope, intercept, r, r2 float64
	nLines := float64(len(spect.Lines))
	for line, peak := range spect.Lines {
		sumX = sumX + peak.Channel
		sumY = sumY + lineMap[line]
		sumXX = sumXX + (peak.Channel)*(peak.Channel)
		sumXY = sumXY + (peak.Channel)*(lineMap[line])
		sumYY = sumYY + (lineMap[line])*(lineMap[line])
	}
	aveX = sumX / nLines
	aveY = sumY / nLines
	Sxx = (sumXX / nLines) - aveX*aveX
	Syy = (sumYY / nLines) - aveY*aveY
	Sxy = (sumXY / nLines) - aveX*aveY
	slope = Sxy / Sxx
	intercept = aveY - slope*aveX
	r = Sxy / (math.Sqrt(Sxx) * math.Sqrt(Syy))
	r2 = r * r
	spect.Gain = slope
	spect.Offset = intercept
	spect.R2 = r2
}
