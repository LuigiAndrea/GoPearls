package fastbinary

// 4.6.8
//fastSearch with fix size array with 10 elements
func fastSearch(a [10]int, v int) (p int) {
	n := len(a)
	i := 8 //largest power of two less than 10, tune depending on the size of the array
	l := -1

	if a[i-1] < v {
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

//fastSearch2 with fix size array with 10 elements, removes overhead of loop and division
func fastSearch2(a [10]int, v int) (p int) {
	n := len(a)
	l := -1
	p = 8

	if a[p-1] < v {
		l = n - p
	}

	p = 4
	if a[l+p] < v {
		l += p
	}

	p = 2
	if a[l+p] < v {
		l += p
	}

	p = 1
	if a[l+p] < v {
		l += p
	}

	p = l + 1
	if p == n || a[p] != v {
		p = -1
	}

	return
}
