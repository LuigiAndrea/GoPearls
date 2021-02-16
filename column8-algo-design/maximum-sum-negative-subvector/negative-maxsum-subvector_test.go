// +build all column8 emptyvector negativemaxsubvector negativevector

package negativemaxsubvector

import (
	"math"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	vector        []float64
	expectedValue float64
}

func TestMaxSubVector(t *testing.T) {

	tests := []testData{
		{vector: []float64{-2.4, 8, 3.3, -5.1, 2, -13}, expectedValue: 11.3},
		{vector: []float64{-2.4, -5.1, -13}, expectedValue: -2.4},
		{vector: []float64{-2.4, -1, -13}, expectedValue: -1},
		{vector: []float64{2.4, -5.1, 13}, expectedValue: 13},
		{vector: []float64{2.4, -1.1, 13}, expectedValue: 14.3},
		{vector: []float64{math.Inf(-1), 1.5, -1.0, 13}, expectedValue: 13.5},
		{vector: []float64{math.Inf(-1), math.Inf(-5), math.Inf(-10)}, expectedValue: math.Inf(-3)},
		{vector: []float64{4, math.Inf(1), 3, 500}, expectedValue: math.Inf(1)},
		{vector: []float64{}, expectedValue: math.Inf(-1)},
		{vector: nil, expectedValue: math.Inf(-1)},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.expectedValue, maxSubvector(test.vector)); err != nil {
			t.Errorf("(%s)%v", m.GetFuncName(maxSubvector), m.ErrorMessageTestCount(i+1, err))
		}
	}
}
