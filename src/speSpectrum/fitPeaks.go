package speSpectrum

import (
	"math"
)

const minStdev float64 = 1.0
const maxStdev float64 = 30.0
const stepStdev float64 = 0.2
const fwhmMult float64 = 2.35482 /// approximation of 2*SQRT(2*LN(2))

func fitPeak(peakPosition []float64, channels []float64) *Peak {
	var stdev, chiSqPrev, chiSqNew float64
	peak := new(Peak)
	peak.Channel = peakPosition[0]

	stdev = minStdev
	chiSqPrev = gaussianChiSq(peakPosition, channels, stdev)
	for stdev <= maxStdev {
		stdev = stdev + stepStdev
		chiSqNew = gaussianChiSq(peakPosition, channels, stdev)
		if chiSqNew > chiSqPrev {
			stdev = stdev - stepStdev
			peak.Area = channels[int(peakPosition[0])] * stdev * math.Sqrt(2*math.Pi)
			peak.FWHM = fwhmMult * stdev
			peak.ChiSq = chiSqPrev
			break
		}
		chiSqPrev = chiSqNew
	}
	return peak
}

func gaussianChiSq(peakPosition []float64, channels []float64, stdev float64) float64 {
	var y, A, x, mu float64
	var chiSq float64
	peakCh := peakPosition[0]
	peakLL := peakPosition[1]
	peakUL := peakPosition[2]
	A = channels[int(peakCh)]
	mu = peakCh
	chiSq = 0
	for i := int(peakLL); i <= int(peakUL); i++ {
		x = float64(i)
		y = gaussianY(x, mu, A, stdev)
		chiSq = chiSq + chiSqY(channels[i], y)
	}
	return chiSq
}

func gaussianY(x float64, mu float64, A float64, stdev float64) float64 {
	expN := -math.Pow(x-mu, 2)
	expD := 2.0 * math.Pow(stdev, 2)
	g := A * math.Exp(expN/expD)
	return g
}

func chiSqY(observed float64, expected float64) float64 {
	chiSq := math.Pow(observed-expected, 2) / expected
	return chiSq
}
