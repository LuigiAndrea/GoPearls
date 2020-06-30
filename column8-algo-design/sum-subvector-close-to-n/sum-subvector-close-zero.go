package sumsubvectornumber

import (
	"math"
	"sort"
)

// Return the subvector with the sum closest to zero,
// if there are more subvectors with the same sum closest to zero return the first occurrence
func sumSubvectorCloseToZero(vector []float64) []float64 {
	var size int

	if size = validateAndReturnSize(vector); size == 0 {
		return []float64{}
	}

	cumArray := calculateCummulativeArray(vector)
	sort.Stable(cumArray)

	currentCloseNumber := newSubvector(0, cumArray[0].originalIndex, math.Abs(cumArray[0].value))

	for i := 0; i < size-1; i++ {
		fromZeroToCurrent := math.Abs(cumArray[i].value)
		if fromZeroToCurrent < currentCloseNumber.distanceToNumber {
			currentCloseNumber = newSubvector(0, cumArray[i].originalIndex, fromZeroToCurrent)
		}

		// Calculate the difference from i and i+1 in the cummulative array
		diff := getDistance(cumArray[i].value, cumArray[i+1].value)
		if diff < currentCloseNumber.distanceToNumber {
			var fromIdx, toIdx int
			if cumArray[i].originalIndex < cumArray[i+1].originalIndex {
				fromIdx, toIdx = cumArray[i].originalIndex, cumArray[i+1].originalIndex
			} else {
				fromIdx, toIdx = cumArray[i+1].originalIndex, cumArray[i].originalIndex
			}

			currentCloseNumber = newSubvector(fromIdx+1, toIdx, diff)
		}
	}

	// Check last element of the cummulative array
	fromZeroToCurrent := math.Abs(cumArray[size-1].value)
	if fromZeroToCurrent < currentCloseNumber.distanceToNumber {
		currentCloseNumber = newSubvector(0, cumArray[size-1].originalIndex, fromZeroToCurrent)
	}

	return vector[currentCloseNumber.from : currentCloseNumber.to+1]
}
