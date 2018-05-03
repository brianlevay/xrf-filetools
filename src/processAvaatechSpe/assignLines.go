package processAvaatechSpe

import ()

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
		peak.Height = spect.Signal[channelInt]
		peak.Total = spect.SPE.Counts[channelInt]
		spect.Lines[secondary_lineList[i].Name] = peak
	}
}
