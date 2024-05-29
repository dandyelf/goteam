package analyser_test

import (
	an "anomaly/analyser"
	"fmt"
)

func ExampleAnomalyAnalyser() {
	fmt.Println(an.AnomalyAnalyser([]float64{1, 12, 5, 23, 34, 35453, 345, 345, 34, 2}))
	// Output: 3625.4 10610.000105560792 <nil>
}
