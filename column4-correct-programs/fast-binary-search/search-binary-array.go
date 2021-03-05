package fastsearch

// 4.6.2
func searchFirstOccurrenceArray(a [100]int, v int) (p int) {
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

// 4.6.8
//fastSearch with fix size array with 100 elements
func fastSearchArray(a [100]int, v int) (p int) {
	n := len(a)
	i := 64 //largest power of two less than 100 (tuned with the length of the array), tune depending on the size of the array
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

//fastSearch2 with fix size array with 100 elements, removes overhead of loop and division
func fastSearch2Array(a [100]int, v int) (p int) {
	n := len(a)
	l := -1
	p = 64

	if a[p-1] < v {
		l = n - p
	}

	p = 32
	if a[l+p] < v {
		l += p
	}

	p = 16
	if a[l+p] < v {
		l += p
	}

	p = 8
	if a[l+p] < v {
		l += p
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
