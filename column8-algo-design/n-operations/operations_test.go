// +build all nop operations

package operations

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	value    float64
	from, to int
}

var op = NewOperation(5)

func TestOperationResetValues(t *testing.T) {

	op.RunOperation(10, 0, 0)

	if err := a.AssertDeepEqual(10.0, op.Values[0]); err != nil {
		t.Error(err)
	}

	op.Reset()

	if err := a.AssertDeepEqual(0.0, op.Values[0]); err != nil {
		t.Error(err)
	}

}

func TestOperations(t *testing.T) {

	tests := []testData{
		testData{value: 50, from: 0, to: 3},
		testData{value: 5, from: 2, to: 3},
		testData{value: -15, from: 1, to: 4},
	}

	for _, test := range tests {
		if err := a.AssertDeepEqual(true, op.RunOperation(test.value, test.from, test.to)); err != nil {
			t.Error(err)
			return
		}
	}

	if err := a.AssertSlicesEqual(a.Float64SlicesMatch{Expected: []float64{50, 35, 40, 40, -15}, Actual: op.Values}); err != nil {
		t.Error(err)
	}

	op.Reset()

	if err := a.AssertSlicesEqual(a.Float64SlicesMatch{Expected: make([]float64, op.Size), Actual: op.Values}); err != nil {
		t.Error(err)
	}

	op2 := NewOperation(2)
	op2.LoadOperation([]float64{2, 3, 5, 10, 12})
	op2.RunOperation(2, 0, 0)
	if err := a.AssertSlicesEqual(a.Float64SlicesMatch{Expected: []float64{4, 3}, Actual: op2.Values}); err != nil {
		t.Error(err)
	}

	op3 := NewOperation(4)
	op3.LoadOperation([]float64{2, 3})
	op3.RunOperation(2, 0, 1)
	if err := a.AssertSlicesEqual(a.Float64SlicesMatch{Expected: []float64{4, 5, 0, 0}, Actual: op3.Values}); err != nil {
		t.Error(err)
	}
}

func TestOperationsWrongIntervals(t *testing.T) {

	tests := []testData{
		testData{value: 5, from: -1, to: 3},
		testData{value: 5, from: 2, to: 1},
		testData{value: 5, from: 1, to: 5},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(false, op.RunOperation(test.value, test.from, test.to)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestOperationWithValues(t *testing.T) {

	op.LoadOperation([]float64{2, 3, 5, 10, 12, 33, 44, 6, 56})
	op.RunOperation(1, 1, 2)
	op.RunOperation(2, 0, 3)
	op.RunOperation(10, 3, 4)

	if err := a.AssertSlicesEqual(a.Float64SlicesMatch{Expected: []float64{4, 6, 8, 22, 22}, Actual: op.Values}); err != nil {
		t.Errorf(err.Error())
	}
}
