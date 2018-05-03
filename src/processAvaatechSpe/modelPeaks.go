package processAvaatechSpe

import (
	"math"
	reg "regression"
)

// Right now, this only gets the maximum intensities of the peaks and the height relative to
// the modelled background underneath. A previous attempt to model the peaks with Gaussian curves
// produced very noisy results and was dropped. I might try again another time.

const linearDist int = 2

func (spect *Spectrum) ModelPeaks(peakCutoff float64) {
	var start, end, p int
	var X, Y []float64
	var linear *reg.Linear
	var prevSlope, currentSlope float64
	var peakList []*Peak
	var peak *Peak

	start = linearDist
	end = spect.MaxChannel - linearDist
	X = make([]float64, 2*linearDist+1)
	for k, _ := range X {
		X[k] = float64(k)
	}
	prevSlope = 0.0
	for i := start; i <= end; i++ {
		Y = spect.Signal[i-linearDist : i+linearDist+1]
		linear = reg.LinearFit(X, Y)
		currentSlope = linear.Slope
		if (prevSlope > 0) && (currentSlope < 0) {
			p = i
			if math.Abs(prevSlope) < math.Abs(currentSlope) {
				p = i - 1
			}
			if spect.Signal[p] >= peakCutoff {
				peak = new(Peak)
				peak.Channel = float64(p)
				peak.Height = spect.Signal[p]
				peak.Total = spect.SPE.Counts[p]
				peakList = append(peakList, peak)
			}
		}
		prevSlope = currentSlope
	}
	spect.Peaks = peakList
}
