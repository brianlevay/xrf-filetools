package processAvaatechSpe

import (
	reg "regression"
)

func (spect *Spectrum) CalculateEnergyScale(gainMidKeV float64) {
	var X, Y []float64
	nLines := float64(len(spect.Lines))
	if nLines >= 2 {
		for line, peak := range spect.Lines {
			X = append(X, peak.Channel)
			Y = append(Y, lineMap[line])
		}
		linear := reg.LinearFit(X, Y)
		spect.Gain = linear.Slope
		spect.Offset = linear.Intercept
		spect.R2 = linear.R2
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
