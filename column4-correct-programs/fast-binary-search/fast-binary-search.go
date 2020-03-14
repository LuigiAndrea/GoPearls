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
	if u == n || a[u] != v {
		p = -1
	}

	return
}
