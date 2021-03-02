package maximumvalue

import "fmt"

//MaxInArray return a the max value in an array
func MaxInArray(array []float64) (float64, error) {
	if len(array) < 1 {
		return 0, fmt.Errorf("Array cannot be empty or nil")
	}

	max := array[0]

	for i := 1; i < len(array); i++ {
		if array[i] > max {
			max = array[i]
		}
	}
	return max, nil
}
