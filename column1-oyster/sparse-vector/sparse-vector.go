package sparsevector

import "fmt"

type sparseVectorInfo struct {
	from, to, data []int
	top            int
}

var sparseInfo sparseVectorInfo

//SparseVector array with sparse elements
type SparseVector struct {
	length int
}

//NewSparseVector create a new sparse vector
func NewSparseVector(sparseLength int) (sv *SparseVector, err error) {
	if sparseLength < 0 {
		return sv, fmt.Errorf("Negative n '%d' argument in SparseVector", sparseLength)
	}

	sparseInfo = sparseVectorInfo{
		from: make([]int, sparseLength),
		to:   make([]int, sparseLength),
		data: make([]int, sparseLength),
		top:  0,
	}

	return &SparseVector{length: len(sparseInfo.data)}, err
}

func (sparse SparseVector) get(index int) (int, error) {
	if sparseInfo.from[index] < sparseInfo.top &&
		index == sparseInfo.to[sparseInfo.from[index]] {
		return sparseInfo.data[index], nil
	}
	return 0, fmt.Errorf("The element at index '%d' has not been initialized", index)
}

func (sparse SparseVector) add(index int, value int) {
	if _, err := sparse.get(index); err == nil {
		sparseInfo.data[index] = value
	} else {
		sparseInfo.data[index] = value
		sparseInfo.from[index] = sparseInfo.top
		sparseInfo.to[sparseInfo.top] = index
		sparseInfo.top++
	}
}
