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

	op.ComputeValues()

	if err := a.AssertSlicesEqual(a.Float64SlicesMatch{Expected: Operation{50, 35, 40, 40, -15}, Actual: op}); err != nil {
		t.Error(err)
	}

	op.Reset()
	expected := make(Operation, 5)
	if err := a.AssertSlicesEqual(a.Float64SlicesMatch{Expected: expected, Actual: op}); err != nil {
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

func TestComputeValuesMultipleTimes(t *testing.T) {

	op.RunOperation(10, 1, 3)
	op.RunOperation(5, 1, 2)
	op.ComputeValues()
	op.RunOperation(10, 0, 1)
	op.ComputeValues()
	op.RunOperation(3, 3, 4)
	op.ComputeValues()

	if err := a.AssertSlicesEqual(a.Float64SlicesMatch{Expected: Operation{10, 25, 15, 13, 3}, Actual: op}); err != nil {
		t.Errorf(err.Error())
	}

	op.Reset()
}

func TestOperationWithValues(t *testing.T) {

	op.LoadOperation([]float64{2, 3, 5, 10, 12})
	op.RunOperation(1, 1, 2)
	op.RunOperation(2, 0, 3)
	op.RunOperation(10, 3, 4)
	op.ComputeValues()

	if err := a.AssertSlicesEqual(a.Float64SlicesMatch{Expected: Operation{4, 6, 8, 22, 22}, Actual: op}); err != nil {
		t.Errorf(err.Error())
	}
}
