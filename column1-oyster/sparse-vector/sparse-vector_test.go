// +build all column1 sparseVector

package sparsevector

import (
	"testing"
)

func TestSparseVector(t *testing.T) {
	sparseVector, _ := NewSparseVector(10)
	expectedValues := map[int]int{3: 51, 2: 50, 9: 12, 8: 0}
	expectedInitializedElement := len(expectedValues)
	var numberOfInitializedElement int

	sparseVector.add(3, 51)
	sparseVector.add(2, 5)
	sparseVector.add(9, 12)
	sparseVector.add(2, 50)
	sparseVector.add(8, 0)

	for i := 0; i < sparseVector.length; i++ {
		if num, err := sparseVector.get(i); err == nil {
			numberOfInitializedElement++
			if num != expectedValues[i] {
				t.Errorf("Expected value %d - Actual value %d", expectedValues[i], num)
			} else {
				t.Logf("(Index: %d, Num: %d)", i, num)
			}
		}
	}

	if numberOfInitializedElement != len(expectedValues) {
		t.Errorf("Expected value %d - Actual value %d", expectedInitializedElement, numberOfInitializedElement)
	} else {
		t.Logf("%d elements have been initialized", numberOfInitializedElement)
	}
}

func TestSparseVectorNegativeIndex(t *testing.T) {
	indexesAdd := []int{10, -4, 3}
	indexesGet := []int{12, 3, -4}

	sparseVector, _ := NewSparseVector(3)
	for _, v := range indexesAdd {
		err := sparseVector.add(v, 100)
		checkError(t, v, err)
	}

	for _, v := range indexesGet {
		_, err := sparseVector.get(v)
		checkError(t, v, err)
	}

	idx := -3
	_, err := NewSparseVector(idx)
	checkError(t, idx, err)
}

func checkError(t *testing.T, idx int, err error) {
	if err == nil {
		t.Errorf("Expected index '%d' out of range", idx)
	} else {
		t.Log(err.Error())
	}
}
