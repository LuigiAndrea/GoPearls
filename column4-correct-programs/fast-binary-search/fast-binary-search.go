package fastsearch

// 4.6.8
//fastSearch with fix size array with 1000 elements
func fastSearch(a []int, v int) (p int) {
	n := len(a)
	i := 512 //largest power of two less than 1000 (tuned with the length of the array), tune depending on the size of the array
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

//fastSearch2 with fix size array with 1000 elements, removes overhead of loop and division
func fastSearch2(a []int, v int) (p int) {
	n := len(a)
	l := -1
	p = 512

	if a[p-1] < v {
		l = n - p
	}

	p = 256
	if a[l+p] < v {
		l += p
	}

	p = 128
	if a[l+p] < v {
		l += p
	}

	p = 64
	if a[l+p] < v {
		l += p
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
