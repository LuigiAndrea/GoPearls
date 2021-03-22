package recurrence

import "fmt"

//an = a[n-1]*c[0] + a[n-2]*c[1] + ... + a[0]*c[k-1] + c[k]
//Return a[0] through a[m-1]
func linearRecurrence(a, c []int, m int) ([]int, error) {
	if err := validateRecurrenceInput(a, c, m); err != nil {
		return nil, err
	}

	result := make([]int, m)
	result[0] = c[0]
	for i := 1; i < m; i++ {
		y := 0
		for j := 0; j < i; j++ {
			y += a[j] * c[i-j-1]
		}
		result[i] = y + c[i]
	}

	return result, nil
}

func validateRecurrenceInput(a, c []int, m int) error {
	sizeA := len(a)
	if sizeA >= len(c) || m-1 > sizeA || m < 1 {
		return fmt.Errorf("The inputs provided are wrong a: %#v c: %#v m: %d", a, c, m)
	}

	return nil
}
