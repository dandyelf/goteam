package analyser

import (
	"math"
)

func AnomalyAnalyser(distribution []float64) (mean float64, stdDev float64, err error) {
	// mean
	sum := 0.0
	for _, value := range distribution {
		sum += value
	}
	mean = sum / float64(len(distribution))
	// stdDev
	variance := 0.0
	for _, value := range distribution {
		variance += math.Pow(value-mean, 2)
	}
	stdDev = math.Sqrt(variance / float64(len(distribution)))
	return
}
