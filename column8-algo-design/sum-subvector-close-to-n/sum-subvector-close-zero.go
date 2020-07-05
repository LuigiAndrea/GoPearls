package sumsubvectornumber

import (
	"math"
	"sort"

	u "github.com/LuigiAndrea/GoPearls/utilities"
)

// Return the subvector with the sum closest to zero,
// if there are more subvectors with the same sum closest to zero return the first occurrence
func sumSubvectorCloseToZero(vector []float64) []float64 {
	var size int

	if size = validateAndReturnSize(vector); size == 0 {
		return []float64{}
	}

	cumArray := u.CalculateCummulativeArray(vector)
	sort.Stable(cumArray)

	currentCloseNumber := newSubvector(0, cumArray[0].OriginalIndex, math.Abs(cumArray[0].Value))

	for i := 0; i < size-1; i++ {
		fromZeroToCurrent := math.Abs(cumArray[i].Value)
		if fromZeroToCurrent < currentCloseNumber.distanceToNumber {
			currentCloseNumber = newSubvector(0, cumArray[i].OriginalIndex, fromZeroToCurrent)
		}

		// Calculate the difference from i and i+1 in the cummulative array
		diff := cumArray.GetDistance(i, i+1)
		if diff < currentCloseNumber.distanceToNumber {
			var fromIdx, toIdx int
			if cumArray[i].OriginalIndex < cumArray[i+1].OriginalIndex {
				fromIdx, toIdx = cumArray[i].OriginalIndex, cumArray[i+1].OriginalIndex
			} else {
				fromIdx, toIdx = cumArray[i+1].OriginalIndex, cumArray[i].OriginalIndex
			}

			currentCloseNumber = newSubvector(fromIdx+1, toIdx, diff)
		}
	}

	// Check last element of the cummulative array
	fromZeroToCurrent := math.Abs(cumArray[size-1].Value)
	if fromZeroToCurrent < currentCloseNumber.distanceToNumber {
		currentCloseNumber = newSubvector(0, cumArray[size-1].OriginalIndex, fromZeroToCurrent)
	}

	return vector[currentCloseNumber.from : currentCloseNumber.to+1]
}
