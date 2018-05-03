package snip

import (
	"math"
)

func RemoveBackground(data []float64, snipWidth float64) []float64 {
	var nPts int
	var signal []float64
	background := fitSpectrumBackground(data, snipWidth)
	nPts = len(background)
	signal = make([]float64, nPts)
	for i := 0; i < nPts; i++ {
		signal[i] = data[i] - background[i]
	}
	return signal
}

func fitSpectrumBackground(data []float64, snipWidth float64) []float64 {
	var vOld, vNew, back []float64
	var nPts, i int
	var snipDistance, currIndex, maxIndex, leftIndex, rightIndex, aveBounds float64
	nPts = len(data)
	maxIndex = float64(nPts - 1)
	vOld = make([]float64, nPts)
	vNew = make([]float64, nPts)
	back = make([]float64, nPts)
	for i = 0; i < nPts; i++ {
		vOld[i] = YtoV(data[i])
	}
	snipDistance = snipWidth / 2
	for snipDistance >= 2 {
		for i = 0; i < nPts; i++ {
			currIndex = float64(i)
			leftIndex = math.Max(0, currIndex-snipDistance)
			rightIndex = math.Min(maxIndex, currIndex+snipDistance)
			aveBounds = (vOld[int(leftIndex)] + vOld[int(rightIndex)]) / 2
			vNew[i] = math.Min(vOld[i], aveBounds)
		}
		snipDistance = snipDistance / 2
		copy(vOld, vNew)
	}
	for i = 0; i < nPts; i++ {
		back[i] = VtoY(vNew[i])
	}
	return back
}

func YtoV(y float64) float64 {
	return math.Log10(math.Log10(math.Sqrt(y+1)+1) + 1)
}

func VtoY(v float64) float64 {
	return math.Pow(math.Pow(10, math.Pow(10, v)-1)-1, 2) - 1
}
