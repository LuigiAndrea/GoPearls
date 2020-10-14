// +build all column1 sparseVector

package sparsevector

import (
	"errors"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

func TestSparseVector(t *testing.T) {
	sparseVector, _ := NewSparseVector(10)
	expectedValues := map[int]int{3: 51, 2: 50, 9: 12, 7: 0, 8: 0}
	expectedInitializedElements := len(expectedValues)
	var numberOfInitializedElements int

	for i, v := range expectedValues {
		sparseVector.add(i, v)
	}

	for i := 0; i < sparseVector.length; i++ {
		if num, err := sparseVector.get(i); err == nil {
			numberOfInitializedElements++
			if err := a.AssertDeepEqual(expectedValues[i], num); err != nil {
				t.Error(m.ErrorMessageTestCount(i, err))
			}
		}
	}

	if err := a.AssertDeepEqual(expectedInitializedElements, numberOfInitializedElements); err != nil {
		t.Error(err)
	}

	//test change value in sparse vector
	indexToChange, newValue := 3, 10
	sparseVector.add(indexToChange, newValue)
	num, _ := sparseVector.get(indexToChange)
	if err := a.AssertDeepEqual(newValue, num); err != nil {
		t.Error(err)
	}
}

func TestSparseVectorNegativeOutRangeIndexGetAndSet(t *testing.T) {
	indexes := []int{12, 3, -4}

	sparseVector, _ := NewSparseVector(3)
	for _, v := range indexes {
		er := sparseVector.add(v, 100)
		checkError(t, er)
	}

	for _, v := range indexes {
		_, er := sparseVector.get(v)
		checkError(t, er)
	}
}

func TestSparseVectorElementNotInitialized(t *testing.T) {
	sparseVector, _ := NewSparseVector(2)
	_, er := sparseVector.get(1)
	checkError(t, er)
}

func TestSparseVectorNegativeSparseVectorSize(t *testing.T) {
	_, er := NewSparseVector(-3)
	checkError(t, er)
}

func checkError(t *testing.T, er error) {
	if err := a.AssertException(errors.New(""), er); err != nil {
		t.Error(err)
	}
}
