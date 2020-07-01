package sumsubvectornumber

import "github.com/LuigiAndrea/GoPearls/utilities"

// Return the subvector with the sum closest to a number
func sumSubvectorCloseToNumber(vector []float64, number float64) []float64 {
	var size int

	if size = validateAndReturnSize(vector); size == 0 {
		return []float64{}
	}

	cumArray := utilities.CalculateCummulativeArray(vector)
	currentCloseNumber := newSubvector(0, 0, getDistance(cumArray[0].Value, number))

	for i := 0; i < size-1; i++ {
		currentNumber := cumArray[i].Value

		// Check i element of the cummulative array
		fromZeroToCurrent := getDistance(cumArray[i].Value, number)
		if fromZeroToCurrent < currentCloseNumber.distanceToNumber {
			currentCloseNumber = newSubvector(0, i, fromZeroToCurrent)
		}

		for j := i + 1; j < size; j++ {
			currentDiff := getDistance(cumArray[j].Value-currentNumber, number)
			if currentDiff < currentCloseNumber.distanceToNumber {
				currentCloseNumber = newSubvector(i+1, j, currentDiff)
			}
		}
	}

	// Check last element of the cummulative array
	fromZeroToCurrent := getDistance(cumArray[size-1].Value, number)
	if fromZeroToCurrent < currentCloseNumber.distanceToNumber {
		currentCloseNumber = newSubvector(0, size-1, fromZeroToCurrent)
	}

	return vector[currentCloseNumber.from : currentCloseNumber.to+1]
}
