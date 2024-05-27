package analyser_test

import (
	an "anomaly/analyser"
	"fmt"
	"math/rand"
	"testing"
)

func TestAnomalyAnalyser(t *testing.T) {
	for _, c := range []struct {
		in []float64
		want []float64
	}{
		{[1, 12, 5], {1.2, 12, }},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	} {
		got := an.AnomalyAnalyser(c.in)
		if got != c.want {
			t.Errorf("String(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
