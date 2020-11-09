// +build all column9 maximumvalue maxnarray

package maximumvalue

import (
	"errors"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	array    []float64
	maxValue float64
	n        int
}

func TestMaximumValueNArray(t *testing.T) {
	tests := []testData{
		testData{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 11.5},
		testData{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 4.54, n: 2},
		testData{array: []float64{5.1}, maxValue: 5.1},
		testData{array: []float64{-3.5, -2.00, -11.5, 9.8, 5.1}, maxValue: -2.00, n: 3},
	}

	for i, test := range tests {
		x, _ := NewX(test.array)
		if test.n > 0 {
			x.SetN(test.n)
		}
		if err := a.AssertDeepEqual(test.maxValue, x.Max()); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestMaximumValueNArrayExceptions(t *testing.T) {

	tests := []testData{
		testData{array: nil, maxValue: 0},
		testData{array: []float64{}, maxValue: 2},
		testData{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 4.54},
		testData{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 4.54, n: -4},
		testData{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 4.54, n: 6},
	}

	for i, test := range tests {
		x, e := NewX(test.array)
		if e == nil {
			e = x.SetN(test.n)
		}

		if err := a.AssertException(errors.New(""), e); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
