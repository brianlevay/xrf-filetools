package speSpectrum

import ()

func fitPeak(peakCh int64, channels []int64) *Peak {
	peak := new(Peak)
	peak.Channel = peakCh
	return peak
}
