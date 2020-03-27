package bsdeclaration

var offset int

type bsdeclaration func([]int, int) int

func (bsd *bsdeclaration) run(A []int, v int) int {
	sizeArray := len(A)
	if sizeArray < 1 {
		defer bsd.setOffset(0)
		return -1
	}

	midPoint := (sizeArray - 1) / 2
	if v == A[midPoint] {
		defer bsd.setOffset(0)
		return offset + midPoint
	} else if v < A[midPoint] {
		return bsd.run(A[:midPoint], v)
	} else {
		bsd.setOffset(offset + midPoint + 1)
		return bsd.run(A[midPoint+1:], v)
	}
}

func (bsd *bsdeclaration) setOffset(os int) {
	offset = os
}
