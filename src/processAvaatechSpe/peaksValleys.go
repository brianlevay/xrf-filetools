package processAvaatechSpe

import (
	"math"
)

const maxChannel int = 500

func getPeaks(inflections [][]float64, peakCutoff float64) []*Peak {
	var peakList []*Peak
	var peak *Peak
	var heightA, heightB, heightAve float64

	nInf := len(inflections)
	for i := 1; i < nInf; i++ {
		if (inflections[i][2] == 1) && (inflections[i-1][2] == -1) &&
			(inflections[i+1][2] == -1) {
			heightA = inflections[i][1] - inflections[i-1][1]
			heightB = inflections[i][1] - inflections[i+1][1]
			heightAve = (heightA + heightB) / 2
			if heightAve >= peakCutoff {
				peak = new(Peak)
				peak.Channel = inflections[i][0]
				peak.HeightAbs = inflections[i][1]
				peak.HeightRel = heightAve
				peakList = append(peakList, peak)
			}
		}
	}
	return peakList
}

func getInflections(counts []float64) [][]float64 {
	var inflections [][]float64
	var delPrev, delNext float64
	var row []float64

	nChannels := int(math.Min(float64(maxChannel), float64(len(counts))))
	row = make([]float64, 3)
	row[0] = 0
	row[1] = counts[0]
	row[2] = -1
	inflections = append(inflections, row)
	for i := 1; i < (nChannels - 1); i++ {
		delPrev = counts[i] - counts[i-1]
		delNext = counts[i+1] - counts[i]
		if (delPrev > 0) && (delNext < 0) {
			row = make([]float64, 3)
			row[0] = float64(i)
			row[1] = counts[i]
			row[2] = 1
			inflections = append(inflections, row)
		} else if (delPrev < 0) && (delNext > 0) {
			row = make([]float64, 3)
			row[0] = float64(i)
			row[1] = counts[i]
			row[2] = -1
			inflections = append(inflections, row)
		}
	}
	row = make([]float64, 3)
	row[0] = float64(len(counts) - 1)
	row[1] = counts[len(counts)-1]
	row[2] = -1
	inflections = append(inflections, row)
	return inflections
}
