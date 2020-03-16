package fastbinary

func fastSearch(a []int, v int) (p int) {
	n := len(a)
	u := n
	l := -1
	for l+1 < u {
		m := (l + u) / 2
		if a[m] < v {
			l = m
		} else {
			u = m
		}
	}

	p = u
	if p == n || a[p] != v {
		p = -1
	}

	return
}

//array with 10 elements
func fastSearch2(a [10]int, v int) (p int) {
	n := len(a)
	i := 8 //largest power of two less than 10, tune depending on the size of the array
	l := -1

	if a[i] < v {
		l = n - i
	}

	for i != 1 {
		i /= 2
		if a[l+i] < v {
			l += i
		}
	}

	p = l + 1
	if p == n || a[p] != v {
		p = -1
	}

	return
}
