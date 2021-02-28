package maximumvalue

import "fmt"

//XSentinel is as slice that will contain a sentinel as last element
type XSentinel []float64

//NewXWithSentinel return a new object of type XSentinel
func NewXWithSentinel(array []float64) (XSentinel, error) {
	if len(array) < 1 {
		return XSentinel{}, fmt.Errorf("Array cannot be empty or nil")
	}
	return append(array, array[0]), nil
}

//Max calculate the maximum value in the slice x[0...n-1], x[n] has the max value
func (x XSentinel) Max() float64 {
	size := len(x) - 1

	for i := 1; i < size; i++ {
		for x[i] < x[size] {
			i++
		}
		x[size] = x[i]
	}

	return x[size]
}
