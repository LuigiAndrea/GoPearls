// +build column9_bench benchmark maxnarray_bench

package maximumvalue

import "testing"

type benchArray struct {
	dscr  string
	input []float64
}

var tests = []benchArray{
	{dscr: "Increasing", input: []float64{-3.5, -3, -2, -1, 1, 2, 3, 4, 5.1, 9.8, 11.5, 14, 15, 40, 50, 60, 70, 80, 90, 100}},
	{dscr: "Decreasing", input: []float64{100, 90, 80, 70, 60, 50, 40, 15, 14, 11.5, 9.8, 5.1, 4, 3, 2, 1, -1, -2, -3, -3.5}},
	{dscr: "Random", input: []float64{50, 2, 1, 40, 14, -1, 90, -2, 15, 5.1, 4, -3.5, 3, 60, 100, 70, 80, 9.8, 11.5, -3}},
}

func BenchmarkMaximumValueRecursive(b *testing.B) {
	for _, test := range tests {
		b.Run(test.dscr, func(b *testing.B) {
			x, _ := NewX(test.input)
			for i := 0; i < b.N; i++ {
				x.Max()
			}
		})
	}
}

func BenchmarkMaximumValueSentinel(b *testing.B) {
	for _, test := range tests {
		b.Run(test.dscr, func(b *testing.B) {
			x, _ := NewXWithSentinel(test.input)
			for i := 0; i < b.N; i++ {
				x.Max()
			}
		})
	}
}
