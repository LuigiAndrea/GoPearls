// +build column9_bench benchmark

package maximumvalue

import "testing"

func BenchmarkMaximumValueIncreasingArray(b *testing.B) {
	x, _ := NewX([]float64{-3.5, -3, -2, -1, 1, 2, 3, 4, 5.1, 9.8, 11.5, 14, 15, 40, 50, 60, 70, 80, 90, 100})
	for i := 0; i < b.N; i++ {
		x.Max()
	}
}

func BenchmarkMaximumValueDecreasingArray(b *testing.B) {
	x, _ := NewX([]float64{100, 90, 80, 70, 60, 50, 40, 15, 14, 11.5, 9.8, 5.1, 4, 3, 2, 1, -1, -2, -3, -3.5})
	for i := 0; i < b.N; i++ {
		x.Max()
	}
}

func BenchmarkMaximumValueArray(b *testing.B) {
	x, _ := NewX([]float64{50, 2, 1, 40, 14, -1, 90, -2, 15, 5.1, 4, -3.5, 3, 60, 100, 70, 80, 9.8, 11.5, -3})
	for i := 0; i < b.N; i++ {
		x.Max()
	}
}
