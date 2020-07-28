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
