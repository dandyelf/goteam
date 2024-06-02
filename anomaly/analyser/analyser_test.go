package analyser_test

import (
	an "anomaly/analyser"
	"fmt"
)

func ExampleAnomalyAnalyser() {
	fmt.Println(an.MeanStdDevCalc([]float64{}))
	fmt.Println(an.MeanStdDevCalc([]float64{1}))
	fmt.Println(an.MeanStdDevCalc([]float64{1, 2}))
	fmt.Println(an.MeanStdDevCalc([]float64{1, 2, 1}))
	fmt.Println(an.MeanStdDevCalc([]float64{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1}))

	// Output: 0 0 <nil>
	// 1 0 <nil>
	// 1.5 0.5 <nil>
	// 1.3333333333333333 0.4714045207910317 <nil>
	// 1.4545454545454546 0.49792959773196915 <nil>
}
