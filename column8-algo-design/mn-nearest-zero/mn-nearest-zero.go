package indexmnclosezero

import (
	sv "github.com/LuigiAndrea/GoPearls/column8-algo-design"
	u "github.com/LuigiAndrea/GoPearls/utilities"
)

// Return the subvector with the sum closest to zero with integer i (0 ≤ i < n – m)
// if there are more subvectors with the same sum closest to zero return the first occurrence
func mnCloseToZero(m, n int, vector []float64) (index int, result []float64) {
	var size int

	if size = sv.ValidateAndReturnSize(vector); size == 0 || m >= n || m < 0 || n > size {
		index = -1
		return
	}

	cumArray := u.CalculateCummulativeArray(vector)
	mn := n - m

	currentCloseNumber := sv.NewSubvector(index, index+m, cumArray[index+m].Value)

	for i := 1; i < mn; i++ {

		diff := cumArray.GetDistance(i-1, i+m)

		if diff < currentCloseNumber.DistanceToNumber {
			currentCloseNumber = sv.NewSubvector(i, i+m, diff)
			index = i
		}
	}

	result = vector[currentCloseNumber.From : currentCloseNumber.To+1]
	return
}
