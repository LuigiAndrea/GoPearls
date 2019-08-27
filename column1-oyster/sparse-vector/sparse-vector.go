package sparsevector

import (
	"fmt"
)

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

func (s *SparseVector) get(index int) (int, error) {
	if err := s.validateIndex(index); err != nil {
		return 0, err
	}

	if sparseInfo.from[index] < sparseInfo.top &&
		index == sparseInfo.to[sparseInfo.from[index]] {
		return sparseInfo.data[index], nil
	}
	return 0, fmt.Errorf("The element at index '%d' has not been initialized", index)
}

func (s *SparseVector) add(index int, value int) error {
	if err := s.validateIndex(index); err != nil {
		return err
	}

	if _, err := s.get(index); err == nil {
		sparseInfo.data[index] = value
	} else {
		sparseInfo.data[index] = value
		sparseInfo.from[index] = sparseInfo.top
		sparseInfo.to[sparseInfo.top] = index
		sparseInfo.top++
	}

	return nil
}

func (s *SparseVector) validateIndex(index int) error {
	if index < 0 || index > s.length {
		return fmt.Errorf("index %d is Out of Range", index)
	}

	return nil
}
