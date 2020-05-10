package maxsubvector

import "github.com/LuigiAndrea/GoPearls/utilities"

func maxSubvector(vector []float64) (maxsofar float64) {
	n := len(vector)
	maxFromZeroToHere := 0.0
	for i := 0; i < n; i++ {
		maxFromZeroToHere = utilities.Max(maxFromZeroToHere+vector[i], 0)
		maxsofar = utilities.Max(maxsofar, maxFromZeroToHere)
	}

	return
}
