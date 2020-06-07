package negativemaxsubvector

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
		if maxFromZeroToHere == 0.0 { //assign max value in a subvector with all negative value
			maxsofar = utilities.Max(maxsofar, maxFromZeroToHere+vector[i])
		} else {
			maxsofar = utilities.Max(maxsofar, maxFromZeroToHere)
		}
	}

	return
}
