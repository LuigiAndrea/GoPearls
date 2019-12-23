// +build all column1 bitvector

package bitvector

import (
	"testing"
)

func TestBitVectorOperations(t *testing.T) {
	bitVectors, _ := NewBitVector(5000000)
	setValues := []int{500000, 500001, 500005}

	for _, v := range setValues {
		bitVectors.Set(v)
	}

	getValues := []int{500000, 500001, 500005, 500002, 0}
	expectedGetValues := []byte{1, 1, 1, 0, 0}

	for i := 0; i < len(getValues); i++ {
		getValue, _ := bitVectors.Get(getValues[i])
		if getValue != expectedGetValues[i] {
			t.Errorf("\nExpected value %d - Actual value %d", expectedGetValues[i], getValue)
		}
	}

	for _, v := range getValues {
		bitVectors.Clear(v)
	}

	for _, v := range getValues {
		clearValue, _ := bitVectors.Get(v)
		if clearValue != 0 {
			t.Errorf("\nExpected value %d - Actual value %d", 0, clearValue)
		}
	}
}

func TestBitVectorGetOperationWrongValue(t *testing.T) {
	testValue := []int{-200, -2, 5024}
	bitVectors, _ := NewBitVector(5000)

	for i, v := range testValue {
		if _, err := bitVectors.Get(v); err == nil {
			t.Errorf("Expected '%T' exception for value '%d'", IndexBitVectorOutOfRange(v), v)
		} else {
			switch err.(type) {
			case IndexBitVectorOutOfRange:
				t.Logf("%d-Get: %v", i, err)
			default:
				t.Errorf("Expected '%T' exception for value '%d'", IndexBitVectorOutOfRange(v), v)
			}
		}
	}
}

func TestBitVectorSetOperationWrongValue(t *testing.T) {
	testValue := []int{-200, -2, 50000}
	bitVectors, _ := NewBitVector(5000)

	for i, v := range testValue {
		if err := bitVectors.Set(v); err == nil {
			t.Errorf("Expected '%T' exception for value '%d'", IndexBitVectorOutOfRange(v), v)
		} else {
			switch err.(type) {
			case IndexBitVectorOutOfRange:
				t.Logf("%d-Set: %v", i, err)
			default:
				t.Errorf("Expected '%T' exception for value '%d'", IndexBitVectorOutOfRange(v), v)
			}
		}
	}
}

func TestBitVectorClearOperationWrongValue(t *testing.T) {
	testValue := []int{-200, -2, 50000}
	bitVectors, _ := NewBitVector(5000)

	for i, v := range testValue {
		if err := bitVectors.Clear(v); err == nil {
			t.Errorf("Expected '%T' exception for value '%d'", IndexBitVectorOutOfRange(v), v)
		} else {
			switch err.(type) {
			case IndexBitVectorOutOfRange:
				t.Logf("%d-Clear: %v", i, err)
			default:
				t.Errorf("Expected '%T' exception for value '%d'", IndexBitVectorOutOfRange(v), v)
			}
		}
	}
}

func TestBitVectorLength(t *testing.T) {
	testValue := -20

	if _, err := NewBitVector(testValue); err == nil {
		t.Errorf("Expected '%T' for value '%d'", NegativeBitVectorSize(testValue), testValue)
	} else {
		switch err.(type) {
		case NegativeBitVectorSize:
			t.Log(err)
		default:
			t.Errorf("Expected '%T' for value '%d'", NegativeBitVectorSize(testValue), testValue)
		}
	}

	testValues := []int{500000, 31, 0, 32, 2000000, 2000031}
	expectedValues := []int{15626, 1, 1, 2, 62501, 62501}

	for i := 0; i < len(testValues); i++ {
		bitVector, _ := NewBitVector(testValues[i])
		lenBitVector := len(*bitVector)
		if lenBitVector != expectedValues[i] {
			t.Errorf("\nExpected value %d - Actual value %d", expectedValues[i], lenBitVector)
		}
	}
}
