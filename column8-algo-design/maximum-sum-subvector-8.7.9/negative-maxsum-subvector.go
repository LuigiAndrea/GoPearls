package negativemaxsubvector

// 8.7.9 We defined the maximum subvector of an array of negative numbers to be zero, the sum of the empty subvector.
// Suppose that we had instead defined the maximum subvector to be the value of the largest element; how would you change the various programs?

import (
	"math"

	"github.com/LuigiAndrea/GoPearls/utilities"
)

func maxSubvector(vector []float64) (maxsofar float64) {

	n := len(vector)
	maxsofar = math.Inf(-1) //instead of 0.0
	maxFromZeroToHere := 0.0

	for i := 0; i < n; i++ {
		maxFromZeroToHere = utilities.Max(maxFromZeroToHere+vector[i], 0.0)
		if maxFromZeroToHere == 0.0 { //assign right max value in a subvector with all negative value
			maxsofar = utilities.Max(maxsofar, maxFromZeroToHere+vector[i])
		} else {
			maxsofar = utilities.Max(maxsofar, maxFromZeroToHere)
		}
	}

	return
}
