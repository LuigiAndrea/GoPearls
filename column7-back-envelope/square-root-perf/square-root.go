package squareroot

import (
	"math"
)

func squareRoot(n int) float64 {
	var fa float64
	for i := 0; i < n; i++ {
		fa = math.Sqrt(10.0)
	}

	return fa
}

func squareRootVariable(n int) float64 {
	var fa float64
	for i := 0; i < n; i++ {
		fa = math.Sqrt(float64(i))
	}

	return fa
}
