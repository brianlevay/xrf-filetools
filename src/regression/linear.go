package regression

import (
	"math"
)

type Linear struct {
	Slope     float64
	Intercept float64
	R2        float64
}

func LinearFit(x []float64, y []float64) *Linear {
	var sumX, sumY, sumXX, sumXY, sumYY float64
	var aveX, aveY, nFlt float64
	var Sxx, Syy, Sxy float64
	var slope, intercept, r, r2 float64
	var linear *Linear

	linear = new(Linear)
	nInt := len(x)
	nFlt = float64(nInt)
	sumX, sumY, sumXX, sumXY = 0, 0, 0, 0
	for i := 0; i < nInt; i++ {
		sumX = sumX + x[i]
		sumY = sumY + y[i]
		sumXX = sumXX + x[i]*x[i]
		sumXY = sumXY + x[i]*y[i]
		sumYY = sumYY + y[i]*y[i]
	}
	aveX = sumX / nFlt
	aveY = sumY / nFlt
	Sxx = (sumXX / nFlt) - aveX*aveX
	Syy = (sumYY / nFlt) - aveY*aveY
	Sxy = (sumXY / nFlt) - aveX*aveY
	slope = Sxy / Sxx
	intercept = aveY - slope*aveX
	r = Sxy / (math.Sqrt(Sxx) * math.Sqrt(Syy))
	r2 = r * r

	linear.Slope = slope
	linear.Intercept = intercept
	linear.R2 = r2
	return linear
}
