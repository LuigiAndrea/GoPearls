// +build all column9 maximumvalue maxnarray

package maximumvalue

import (
	"errors"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testDataSentinel struct {
	array    []float64
	maxValue float64
}

func TestMaximumValueSentinelArray(t *testing.T) {
	tests := []testDataSentinel{
		{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 11.5},
		{array: []float64{3.5, 4.54, 11.5, 9.8, 5.1, 11.55}, maxValue: 11.55},
		{array: []float64{20, 3.5, 4.54, 11.5, 9.8, 5.1}, maxValue: 20},
		{array: []float64{5.1}, maxValue: 5.1},
		{array: []float64{-3.5, -2.00, -11.5, -9.8, -5.1}, maxValue: -2.00},
	}

	for i, test := range tests {
		x, _ := NewXWithSentinel(test.array)

		if err := a.AssertDeepEqual(test.maxValue, x.Max()); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestMaximumValueSentinelArrayExceptions(t *testing.T) {

	tests := []testData{
		{array: nil, maxValue: 0},
		{array: []float64{}, maxValue: 0},
	}

	for i, test := range tests {
		_, e := NewX(test.array)

		if err := a.AssertException(errors.New(""), e); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
