package fastbinary

// 4.6.2
func searchFirstOccurrence(a []int, v int) (p int) {
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
