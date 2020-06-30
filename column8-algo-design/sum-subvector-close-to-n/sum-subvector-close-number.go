package sumsubvectornumber

// Return the subvector with the sum closest to a number
func sumSubvectorCloseToNumber(vector []float64, number float64) []float64 {
	var size int

	if size = validateAndReturnSize(vector); size == 0 {
		return []float64{}
	}

	cumArray := calculateCummulativeArray(vector)
	currentCloseNumber := newSubvector(0, 0, getDistance(cumArray[0].value, number))

	for i := 0; i < size-1; i++ {
		currentNumber := cumArray[i].value

		fromZeroToCurrent := getDistance(cumArray[i].value, number)
		if fromZeroToCurrent < currentCloseNumber.distanceToNumber {
			currentCloseNumber = newSubvector(0, i, fromZeroToCurrent)
		}

		for j := i + 1; j < size; j++ {
			currentDiff := getDistance(cumArray[j].value-currentNumber, number)
			if currentDiff < currentCloseNumber.distanceToNumber {
				currentCloseNumber = newSubvector(i+1, j, currentDiff)
			}
		}
	}

	// Check last element of the cummulative array
	fromZeroToCurrent := getDistance(cumArray[size-1].value, number)
	if fromZeroToCurrent < currentCloseNumber.distanceToNumber {
		currentCloseNumber = newSubvector(0, size-1, fromZeroToCurrent)
	}

	return vector[currentCloseNumber.from : currentCloseNumber.to+1]
}
