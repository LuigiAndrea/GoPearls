package maxsubvector

import "github.com/LuigiAndrea/GoPearls/utilities"

func maxSubvectorCubic(vector []float64) (maxsofar float64) {
	n := len(vector)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			sum := 0.0
			for k := i; k <= j; k++ {
				sum += vector[k]
			}
			maxsofar = utilities.Max(sum, maxsofar)
		}
	}

	return
}
