package maxsubvector

import (
	"github.com/LuigiAndrea/GoPearls/utilities"
)

func maxSubVectorCrossing(vector []float64) float64 {
	return maxSubvectorDivideAndConquer(vector, 0, len(vector)-1)
}

func maxSubvectorDivideAndConquer(vector []float64, l, u int) float64 {
	if l > u { //zero elements
		return 0
	}

	if l == u { //one element
		return utilities.Max(vector[l], 0)
	}

	m := (l + u) / 2
	lmax, rmax, sum := 0.0, 0.0, 0.0

	//max crossing on the left
	for i := m; i >= l; i-- {
		sum += vector[i]
		lmax = utilities.Max(lmax, sum)
	}

	sum = 0.0
	//max crossing on the right
	for i := m + 1; i <= u; i++ {
		sum += vector[i]
		rmax = utilities.Max(rmax, sum)
	}

	return utilities.Max(lmax+rmax, maxSubvectorDivideAndConquer(vector, l, m), maxSubvectorDivideAndConquer(vector, m+1, u))

}
