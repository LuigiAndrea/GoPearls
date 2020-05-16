package maxsubvector

import (
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
		cumSum[i], _ = utilities.Round(vector[i]+cumSum[i-1], 2)
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		sum = cumSum[i]
		maxsofar = utilities.Max(sum, maxsofar)
	}

	for i := 1; i < n; i++ {
		sum = 0.0
		for j := i; j < n; j++ {
			sum = cumSum[j] - cumSum[i-1]
			maxsofar = utilities.Max(sum, maxsofar)
		}
	}

	return
}
