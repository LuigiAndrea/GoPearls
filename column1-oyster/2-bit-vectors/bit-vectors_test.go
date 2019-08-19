// +build all column1 bitvector

package column1

import (
	"testing"
)

func TestBitVectorOperations(t *testing.T) {
	bitVectors, _ := newBitVector(5000000)
	setValues := []int{500000, 500001, 500005}

	for _, v := range setValues {
		bitVectors.set(v)
	}

	getValues := []int{500000, 500001, 500005, 500002, 0}
	expectedGetValues := []byte{1, 1, 1, 0, 0}

	for i := 0; i < len(getValues); i++ {
		getValue, _ := bitVectors.get(getValues[i])
		if getValue != expectedGetValues[i] {
			t.Errorf("\nExpected value %d - Actual value %d", expectedGetValues[i], getValue)
		}
	}

	for _, v := range getValues {
		bitVectors.clear(v)
	}

	for _, v := range getValues {
		clearValue, _ := bitVectors.get(v)
		if clearValue != 0 {
			t.Errorf("\nExpected value %d - Actual value %d", 0, clearValue)
		}
	}
}

func TestBitVectorGetOperationWrongValue(t *testing.T) {
	testValue := []int{-200, -2, 50000}
	bitVectors, _ := newBitVector(5000)

	for i, v := range testValue {
		if _, err := bitVectors.get(v); err == nil {
			t.Errorf("Expected an Out of Range error")
		} else {
			t.Logf("%d-Get: %v", i, err.Error())
		}
	}
}

func TestBitVectorSetOperationWrongValue(t *testing.T) {
	testValue := []int{-200, -2, 50000}
	bitVectors, _ := newBitVector(5000)

	for i, v := range testValue {
		if err := bitVectors.set(v); err == nil {
			t.Errorf("Expected an Out of Range error")
		} else {
			t.Logf("%d-Set: %v", i, err.Error())
		}

		if err := bitVectors.clear(v); err == nil {
			t.Errorf("Expected an Out of Range error")
		} else {
			t.Logf("%d-Clear: %v", i, err.Error())
		}
	}
}

func TestBitVectorLength(t *testing.T) {

	testValues := []int{500000, 31, 0, 32, 2000000, 2000031}
	expectedValues := []int{15626, 1, 1, 2, 62501, 62501}

	if _, err := newBitVector(-20); err == nil {
		t.Errorf("\nExpected an expection: negative n argument in newBitVector")
	} else {
		t.Logf(err.Error())
	}

	for i := 0; i < len(testValues); i++ {
		bitVector, _ := newBitVector(testValues[i])
		lenBitVector := len(*bitVector)
		if lenBitVector != expectedValues[i] {
			t.Errorf("\nExpected value %d - Actual value %d", expectedValues[i], lenBitVector)
		}
	}
}
