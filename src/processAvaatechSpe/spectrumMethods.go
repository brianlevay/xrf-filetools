package processAvaatechSpe

import (
	"math"
)

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

func (spect *Spectrum) AssignPrimaryLines(gainMinKeV float64, gainMaxKeV float64) {
	var lineMap map[string]*Peak
	var peak *Peak
	var startSearch int
	var maxLineCh, minLineCh float64
	lineMap = make(map[string]*Peak)
	nPeaks := len(spect.Peaks)
	nLines := len(primary_lineList)
	startSearch = 0
	for i := 0; i < nPeaks; i++ {
		peak = spect.Peaks[i]
		for j := startSearch; j < nLines; j++ {
			minLineCh = primary_lineList[j].Energy / gainMaxKeV // maxGain => minCh
			maxLineCh = primary_lineList[j].Energy / gainMinKeV
			if (peak.Channel >= minLineCh) && (peak.Channel <= maxLineCh) {
				lineMap[primary_lineList[j].Name] = peak
				startSearch = j + 1
				break
			}
		}
	}
	spect.Lines = lineMap
}

func (spect *Spectrum) CalculateEnergyScale(gainMidKeV float64) {
	var sumX, sumY, sumXX, sumXY, sumYY float64
	var aveX, aveY float64
	var Sxx, Syy, Sxy float64
	var slope, intercept, r, r2 float64
	sumX = 0
	sumY = 0
	sumXX = 0
	sumXY = 0
	sumYY = 0
	nLines := float64(len(spect.Lines))
	if nLines >= 2 {
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
	} else if nLines == 1 {
		for line, peak := range spect.Lines {
			spect.Gain = lineMap[line] / peak.Channel
			spect.Offset = 0.0
			spect.R2 = 1.00
		}
	} else {
		spect.Gain = gainMidKeV
		spect.Offset = 0.0
		spect.R2 = 1.00
	}
}

func (spect *Spectrum) AssignSecondaryLines() {
	var channelFlt float64
	var channelInt int
	var peak *Peak
	nLines := len(secondary_lineList)
	for i := 0; i < nLines; i++ {
		channelFlt = keVtoChannel(secondary_lineList[i].Energy, spect.Gain, spect.Offset)
		channelInt = int(channelFlt)
		peak = new(Peak)
		peak.Channel = channelFlt
		peak.Height = spect.SPE.Counts[channelInt]
		spect.Lines[secondary_lineList[i].Name] = peak
	}
}
