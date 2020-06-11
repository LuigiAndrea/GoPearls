package sumsubvectornumber

import (
	"math"
	"sort"

	"github.com/LuigiAndrea/GoPearls/utilities"
)

type cummulativeArray []cummulativeValue

func (ca cummulativeArray) Len() int           { return len(ca) }
func (ca cummulativeArray) Less(i, j int) bool { return ca[i].value < ca[j].value }
func (ca cummulativeArray) Swap(i, j int)      { ca[i], ca[j] = ca[j], ca[i] }

type cummulativeValue struct {
	value         float64
	originalIndex int
}

type subvector struct {
	from, to         int
	distanceToNumber float64
}

//subvector constructor to fill currentCloseNumber
func newSubvector(from, to int, value float64) subvector {
	return subvector{
		from:             from,
		to:               to,
		distanceToNumber: math.Abs(value)}
}

func sumSubvectorCloseToZero(vector []float64) []float64 {
	size := len(vector)
	cumArray := calculateCummulativeArray(vector)
	sort.Stable(cumArray)

	currentCloseNumber := newSubvector(0, cumArray[0].originalIndex, cumArray[0].value)

	for i := 0; i < size-1; i++ {
		if math.Abs(cumArray[i].value) < currentCloseNumber.distanceToNumber {
			currentCloseNumber = newSubvector(0, cumArray[i].originalIndex, cumArray[i].value)
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
	if math.Abs(cumArray[size-1].value) < currentCloseNumber.distanceToNumber {
		currentCloseNumber = newSubvector(0, cumArray[size-1].originalIndex+1, cumArray[size-1].value)
	}

	return vector[currentCloseNumber.from : currentCloseNumber.to+1]
}

func calculateCummulativeArray(vector []float64) cummulativeArray {
	size := len(vector)
	cArray := make(cummulativeArray, size)
	sum := 0.0

	for i := 0; i < size; i++ {
		sum, _ = utilities.Round(sum+vector[i], 2)
		cArray[i] = cummulativeValue{value: sum, originalIndex: i}
	}

	return cArray
}

func getDistance(value1, value2 float64) float64 {
	return math.Abs(value1 - value2)
}
