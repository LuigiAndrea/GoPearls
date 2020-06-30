package sumsubvectornumber

import (
	"math"

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
	roundValue, _ := utilities.Round(value, 2)
	return subvector{
		from:             from,
		to:               to,
		distanceToNumber: roundValue}
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

func getDistance(value1, value2 float64) (returnValue float64) {
	returnValue, _ = utilities.Round(math.Abs(value1-value2), 2)
	return
}

func validateAndReturnSize(vector []float64) int {
	var size int

	if vector == nil {
		return size
	}
	if size = len(vector); size == 0 {
		return size
	}

	return size
}
