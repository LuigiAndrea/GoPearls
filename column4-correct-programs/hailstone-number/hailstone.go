package hailstone

func hailstoneSequence(x int) []int {
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

func hailstoneRecursive(x int, result []int) []int {
	if x == 1 {
		return result
	}

	if even(x) {
		x /= 2
		result = append(result, x)
		return hailstoneRecursive(x, result)
	}

	x = x*3 + 1
	result = append(result, x)
	return hailstoneRecursive(x, result)
}

func even(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}
