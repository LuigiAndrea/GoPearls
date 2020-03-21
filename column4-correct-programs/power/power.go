package power

func power(n, e int) int {
	if e < 0 {
		return -1
	}

	return exp(n, e)
}

func exp(n, e int) int {
	if e == 0 {
		return 1
	} else if e%2 == 0 {
		return square(exp(n, e/2))
	}
	e--
	return n * exp(n, e)
}

func square(n int) int {
	return n * n
}
