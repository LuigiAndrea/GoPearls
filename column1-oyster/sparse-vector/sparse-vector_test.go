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
	sparseVector, _ := NewSparseVector(3)
	err := sparseVector.add(10, 3)
	checkError(t, err)
	err = sparseVector.add(-4, 30)
	checkError(t, err)
	_, err = sparseVector.get(12)
	checkError(t, err)
	_, err = sparseVector.get(-6)
	checkError(t, err)
}

func checkError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected an index out of range")
	} else {
		t.Log(err.Error())
	}
}
