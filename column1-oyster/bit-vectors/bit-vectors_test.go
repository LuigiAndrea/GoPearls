// +build all column1 bitvector

package bitvector

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	in, out int
}

type testOperationData struct {
	in  int
	out byte
}

func TestBitVectorOperations(t *testing.T) {
	bitVectors, _ := NewBitVector(5000000)
	setValues := []int{500000, 500001, 500005}

	for _, v := range setValues {
		bitVectors.Set(v)
	}

	tests := []testOperationData{
		testOperationData{in: 500000, out: 1},
		testOperationData{in: 500001, out: 1},
		testOperationData{in: 500005, out: 1},
		testOperationData{in: 500002, out: 0},
		testOperationData{in: 0, out: 0}}

	for i, test := range tests {
		if value, e := bitVectors.Get(test.in); e != nil {
			t.Error(m.ErrorMessageTestCount(i+1, e))
		} else if err := a.AssertDeepEqual(test.out, value); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}

	for _, test := range tests {
		bitVectors.Clear(test.in)
	}

	for i, test := range tests {
		if clearValue, e := bitVectors.Get(test.in); e != nil {
			t.Error(m.ErrorMessageTestCount(i+1, e))
		} else if err := a.AssertDeepEqual(byte(0), clearValue); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestBitVectorSetAndClearOperationsWrongValues(t *testing.T) {
	testValues := []int{-200, -2, 5024}
	bitVectors, _ := NewBitVector(5000)
	funcToTest := []func(int) error{bitVectors.Set, bitVectors.Clear}

	for _, vectorOp := range funcToTest {
		for i, v := range testValues {
			e := vectorOp(v)
			if err := a.AssertDeepException(&IndexBitVectorOutOfRange{size: v}, e); err != nil {
				t.Error(m.ErrorMessageTestCount(i+1, err))
			}
		}
	}
}

func TestBitVectorGetOperationWrongValue(t *testing.T) {
	testValues := []int{-200, -2, 5024}
	bitVectors, _ := NewBitVector(5000)

	for i, v := range testValues {
		_, e := bitVectors.Get(v)
		if err := a.AssertDeepException(&IndexBitVectorOutOfRange{size: v}, e); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestBitVectorLength(t *testing.T) {

	testValue := -20
	_, e := NewBitVector(testValue)
	if err := a.AssertDeepException(&NegativeBitVectorSize{size: testValue}, e); err != nil {
		t.Error(err)
	}

	tests := []testData{
		testData{in: 500000, out: 15626},
		testData{in: 31, out: 1},
		testData{in: 0, out: 1},
		testData{in: 32, out: 2},
		testData{in: 2000000, out: 62501},
		testData{in: 2000031, out: 62501}}

	for i, test := range tests {
		bitVector, _ := NewBitVector(test.in)
		lenBitVector := len(*bitVector)
		if err := a.AssertDeepEqual(test.out, lenBitVector); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
