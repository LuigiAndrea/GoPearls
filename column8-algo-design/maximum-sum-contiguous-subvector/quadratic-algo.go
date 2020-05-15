package maxsubvector

import (
	"math"

	"github.com/LuigiAndrea/GoPearls/utilities"
)

func maxSubvectorQuadratic1(vector []float64) (maxsofar float64) {
	n := len(vector)
	for i := 0; i < n; i++ {
		sum := 0.0
		for j := i; j < n; j++ {
			sum += vector[j]
			maxsofar = utilities.Max(sum, maxsofar)
		}
	}

	return
}

func maxSubvectorQuadratic2(vector []float64) (maxsofar float64) {
	n := len(vector)
	if n == 0 {
		return
	}

	cumSum := make([]float64, n)
	cumSum[0] = vector[0]
	for i := 1; i < n; i++ {
		cumSum[i] = math.Round((vector[i]+cumSum[i-1])*100) / 100
	}

	for i := 0; i < n; i++ {
		sum := 0.0
		for j := i; j < n; j++ {
			if i == 0 {
				sum = cumSum[j]
			} else {
				sum = cumSum[j] - cumSum[i-1]
			}
			maxsofar = utilities.Max(sum, maxsofar)
		}
	}

	return
}
