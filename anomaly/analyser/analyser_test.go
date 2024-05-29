package analyser_test

import (
	an "anomaly/analyser"
	"fmt"
)

func ExampleAnomalyAnalyser() {
	fmt.Println(an.AnomalyAnalyser([]float64{1, 12, 5, 23, 34, 35453, 345, 345, 34, 2}))
	// Output: Mean: 3625.400000
	// StdDev: 10610.000106
	// 3625.4 10610.000105560792 <nil>
}
