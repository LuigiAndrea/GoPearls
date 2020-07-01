package sumsubvectornumber

import (
	"math"

	"github.com/LuigiAndrea/GoPearls/utilities"
)

type subvector struct {
	from, to         int
	distanceToNumber float64
}

//subvector constructor to fill currentCloseNumber
func newSubvector(from, to int, value float64) subvector {
	return subvector{
		from:             from,
		to:               to,
		distanceToNumber: value}
}

//getDistance calculate the difference between two values
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
