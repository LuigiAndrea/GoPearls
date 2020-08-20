package sumsubvectornumber

import (
	sv "github.com/LuigiAndrea/GoPearls/column8-algo-design"
	u "github.com/LuigiAndrea/GoPearls/utilities"
)

// Return the subvector with the sum closest to a number
func sumSubvectorCloseToNumber(vector []float64, number float64) []float64 {
	var size int

	if size = sv.ValidateAndReturnSize(vector); size == 0 {
		return []float64{}
	}

	cumArray := u.CalculateCummulativeArray(vector)
	currentCloseNumber := sv.NewSubvector(0, 0, u.GetDistance(cumArray[0].Value, number))

	for i := 0; i < size-1; i++ {
		currentNumber := cumArray[i].Value

		// Check i element of the cummulative array
		fromZeroToCurrent := u.GetDistance(cumArray[i].Value, number)
		if fromZeroToCurrent < currentCloseNumber.DistanceToNumber {
			currentCloseNumber = sv.NewSubvector(0, i, fromZeroToCurrent)
		}

		for j := i + 1; j < size; j++ {
			currentDiff := u.GetDistance(cumArray[j].Value-currentNumber, number)
			if currentDiff < currentCloseNumber.DistanceToNumber {
				currentCloseNumber = sv.NewSubvector(i+1, j, currentDiff)
			}
		}
	}

	// Check last element of the cummulative array
	fromZeroToCurrent := u.GetDistance(cumArray[size-1].Value, number)
	if fromZeroToCurrent < currentCloseNumber.DistanceToNumber {
		currentCloseNumber = sv.NewSubvector(0, size-1, fromZeroToCurrent)
	}

	return vector[currentCloseNumber.From : currentCloseNumber.To+1]
}
