package speSpectrum

import ()

func (spect *Spectrum) GetPeaks() {
	var peaks [][]int64
	var baseA, heightA, baseB, heightB int64
	var row []int64

	inflections := getInflections(spect.SPE.Counts)
	nInf := len(inflections)
	for i := 1; i < nInf; i++ {
		if (inflections[i][2] == 1) && (inflections[i-1][2] == -1) &&
			(inflections[i+1][2] == -1) {
			baseA = inflections[i-1][1]
			heightA = inflections[i][1] - baseA
			baseB = inflections[i+1][1]
			heightB = inflections[i][1] - baseB
			if (heightA >= peakMin) || (heightB >= peakMin) {
				row = make([]int64, 2)
				row[0] = inflections[i][0]
				row[1] = inflections[i][1]
				peaks = append(peaks, row)
			}
		}
	}
	spect.Peaks = peaks
}

func getInflections(counts []int64) [][]int64 {
	var inflections [][]int64
	var delPrev, delNext int64
	var row []int64

	nChannels := len(counts)
	for i := 1; i < (nChannels - 1); i++ {
		delPrev = counts[i] - counts[i-1]
		delNext = counts[i+1] - counts[i]
		if (delPrev > 0) && (delNext < 0) {
			row = make([]int64, 3)
			row[0] = int64(i)
			row[1] = counts[i]
			row[2] = 1
			inflections = append(inflections, row)
		} else if (delPrev < 0) && (delNext > 0) {
			row = make([]int64, 3)
			row[0] = int64(i)
			row[1] = counts[i]
			row[2] = -1
			inflections = append(inflections, row)
		}
	}
	return inflections
}
