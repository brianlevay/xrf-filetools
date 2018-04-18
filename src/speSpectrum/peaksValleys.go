package speSpectrum

import (
	"math"
)

const maxChannel int = 500
const peakCutoff float64 = 1000.0

func getPeakPositions(inflections [][]float64) [][]float64 {
	var peakPositions [][]float64
	var row []float64
	var baseA, heightA, baseB, heightB float64
	var heightAve float64

	nInf := len(inflections)
	for i := 1; i < nInf; i++ {
		if (inflections[i][2] == 1) && (inflections[i-1][2] == -1) &&
			(inflections[i+1][2] == -1) {
			baseA = inflections[i-1][1]
			heightA = inflections[i][1] - baseA
			baseB = inflections[i+1][1]
			heightB = inflections[i][1] - baseB
			heightAve = (heightA + heightB) / 2
			if heightAve >= peakCutoff {
				row = make([]float64, 3)
				row[0] = inflections[i][0]
				row[1] = inflections[i-1][0]
				row[2] = inflections[i+1][0]
				peakPositions = append(peakPositions, row)
			}
		}
	}
	return peakPositions
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
