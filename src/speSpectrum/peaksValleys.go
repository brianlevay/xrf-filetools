package speSpectrum

import ()

func (spect *Spectrum) GetPeaks(peakMin int64) {
	var peakList []*Peak
	var peak *Peak
	var baseA, heightA, baseB, heightB int64
	var heightAve int64

	inflections := getInflections(spect.SPE.Counts)
	nInf := len(inflections)
	for i := 1; i < nInf; i++ {
		if (inflections[i][2] == 1) && (inflections[i-1][2] == -1) &&
			(inflections[i+1][2] == -1) {
			baseA = inflections[i-1][1]
			heightA = inflections[i][1] - baseA
			baseB = inflections[i+1][1]
			heightB = inflections[i][1] - baseB
			heightAve = int64((float64(heightA) + float64(heightB)) / 2)
			if heightAve >= peakMin {
				peak = new(Peak)
				peak.Channel = inflections[i][0]
				peak.Height = heightAve
				peak.Width = inflections[i+1][0] - inflections[i-1][0]
				peakList = append(peakList, peak)
			}
		}
	}
	spect.Peaks = peakList
}

func getInflections(counts []int64) [][]int64 {
	var inflections [][]int64
	var delPrev, delNext int64
	var row []int64

	nChannels := len(counts)
	row = make([]int64, 3)
	row[0] = int64(0)
	row[1] = counts[0]
	row[2] = -1
	inflections = append(inflections, row)
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
	row = make([]int64, 3)
	row[0] = int64(len(counts) - 1)
	row[1] = counts[len(counts)-1]
	row[2] = -1
	inflections = append(inflections, row)
	return inflections
}
