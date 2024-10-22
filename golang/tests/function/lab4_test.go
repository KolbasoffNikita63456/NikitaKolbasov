package function_test

import (
	"math"
	"testing"

	function "isuct.ru/informatics2022/function"
)

type Test struct {
	x, a, b float64
	out     float64
}

var tests = []Test{
	{0.33, 0.06, 0.05, 1.527163095495038},
	{1.05, 0.06, 0.05, math.NaN()},
	{0.95, 0.06, 0.05, 1.0078412577756953},
	{0.24, 0.05, 0.07, 1.633123935319537e+16},
}

func TestCalculateYWithNaN(t *testing.T) {
	for _, test := range tests {
		if math.IsNaN(test.out) {
			got := function.CalculateY(test.x, test.a, test.b)
			if !math.IsNaN(got) {
				t.Errorf("Test failed for NaN case x=%f, a=%f, b=%f: got %f, want NaN", test.x, test.a, test.b, got)
			}
		}
	}
}

func TestCalculateYWithoutNaN(t *testing.T) {
	for _, test := range tests {
		if !math.IsNaN(test.out) {
			got := function.CalculateY(test.x, test.a, test.b)
			if got != test.out {
				t.Errorf("Test failed for x=%f, a=%f, b=%f: got %f, want %f", test.x, test.a, test.b, got, test.out)
			}
		}
	}
}
