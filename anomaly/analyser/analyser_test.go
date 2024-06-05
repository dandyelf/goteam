package analyser_test

import (
	an "anomaly/analyser"
	"fmt"
)

func ExampleAnomalyAnalyser() {
	data := []float64{1.1, 2.1}
	fmt.Println(an.AnomalyAnalise(1, 0.3))

	// Output: 0 0 <nil>
	// 1 0 <nil>
	// 1.5 0.5 <nil>
	// 1.3333333333333333 0.4714045207910317 <nil>
	// 1.4545454545454546 0.49792959773196915 <nil>
}
