package terminate

func terminate3xplus1(x int) []int {
	var result []int

	for i := 0; x != 1; i++ {
		if even(x) {
			x /= 2
			result = append(result, x)
		} else {
			x = x*3 + 1
			result = append(result, x)
		}
	}

	return result
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
