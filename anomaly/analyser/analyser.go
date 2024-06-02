package analyser

import (
	"log"
	"math"
)

var (
	Mean         float64
	StdDev       float64
	Count        int
	Distribution []float64
)

const (
	MaxDistribution = 95
	PrintCheckpoint = 10
)

func AnomalyAnalise(value float64, STDAnomalyCoefficient float64) (err error) {
	if Count >= MaxDistribution {
		Count++
		if IsAnomaly(value, STDAnomalyCoefficient) {
			log.Println("Anomaly found: ", value, "Count: ", Count)
		}

	} else if Count < MaxDistribution-1 {
		Distribution = append(Distribution, value)
		Count++
	} else {
		Distribution = append(Distribution, value)
		Count++
		meanStdDevCalc()
	}

	return
}

func meanStdDevCalc() (err error) {
	log.Println("MeanStdDev. Count: ", Count)
	if len(Distribution) == 0 {
		log.Println("no stream found")
	}
	// mean
	sum := 0.0
	for _, value := range Distribution {
		sum += value
	}
	Mean = sum / float64(len(Distribution))
	// stdDev
	variance := 0.0
	for _, value := range Distribution {
		variance += math.Pow(value-Mean, 2)
	}
	StdDev = math.Sqrt(variance / float64(len(Distribution)))
	Count = len(Distribution)
	if Count%PrintCheckpoint == 0 {
		log.Println(len(Distribution), "value: ", Distribution[len(Distribution)])
		log.Println("Mean: ", Mean, "STDdev: ", StdDev)
	}
	return
}

func IsAnomaly(value float64, coeff float64) bool {
	stdDevMultiplier := coeff * StdDev
	return math.Abs(value-Mean) > stdDevMultiplier
}
