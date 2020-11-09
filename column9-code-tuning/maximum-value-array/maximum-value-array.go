package maximumvalue

import "fmt"

//X object that contains a slice and its length
type X struct {
	array []float64
	n     int
	size  int
}

//NewX return a new object of type X, default value for n is the slice length
func NewX(xArray []float64) (X, error) {
	size := len(xArray)
	if size < 1 {
		return X{}, fmt.Errorf("Array cannot be empty or nil")
	}

	return X{array: xArray, n: size, size: size}, nil
}

//SetN set the value for n
func (x *X) SetN(n int) error {
	if n < 1 || n > x.size {
		return fmt.Errorf("Value 'n:%d' must be greater than zero and smaller than slice length '%d'", n, x.size)
	}
	x.n = n
	return nil
}

//Max calculate the maximum value in the slice x[0...n-1]
func (x X) Max() float64 {
	return x.arraymax(x.n)
}

func (x X) arraymax(n int) float64 {
	if n == 1 {
		return x.array[0]
	}

	return max(x.array[n-1], x.arraymax(n-1))
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
