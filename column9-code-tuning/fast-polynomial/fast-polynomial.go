package fastpoly

func EvalPolynomial(a []int, x int) int {
	s := len(a)
	xi := 1
	y := 0
	if s > 0 {
		y = a[0]
	}
	for i := 1; i < s; i++ {
		xi = x * xi
		y += xi * a[i]
	}
	return y
}

func EvalPolynomialFast(a []int, x int) int {
	s := len(a)
	y := 0
	if s > 0 {
		y = a[s-1]
	}

	for i := s - 2; i >= 0; i-- {
		y = x*y + a[i]
	}

	return y
}
