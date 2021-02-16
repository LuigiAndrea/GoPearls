// +build all column8 mnclosezero

package indexmnclosezero

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	m, n, index       int
	vector, subvector []float64
}

func TestNearestZero(t *testing.T) {
	tests := []testData{
		{m: 2, n: 4, vector: []float64{1, 4, -1, 2}, index: 0, subvector: []float64{1, 4, -1}},
		{m: 1, n: 4, vector: []float64{1, 4, -1, 2}, index: 2, subvector: []float64{-1, 2}},
		{m: 0, n: 4, vector: []float64{1, 4, -1, 2}, index: 0, subvector: []float64{1}},
		{m: 3, n: 4, vector: []float64{1, 4, -1, 2}, index: 0, subvector: []float64{1, 4, -1, 2}},
		{m: 1, n: 4, vector: []float64{1, 4, -1, 2, 0, 20}, index: 2, subvector: []float64{-1, 2}},
		{m: 2, n: 3, vector: []float64{1, 4, -1, 2}, index: 0, subvector: []float64{1, 4, -1}},
		{m: 2, n: 5, vector: []float64{1, 4, -1, 2}, index: -1, subvector: nil},
		{m: 4, n: 4, vector: []float64{1, 4, -1, 2, 0, 20}, index: -1, subvector: nil},
		{m: 1, n: 4, vector: []float64{}, index: -1, subvector: nil},
		{m: -1, n: 4, vector: []float64{}, index: -1, subvector: nil},
		{m: 1, n: 0, vector: []float64{}, index: -1, subvector: nil},
	}

	for i, test := range tests {
		idx, result := mnCloseToZero(test.m, test.n, test.vector)

		if err := a.AssertDeepEqual(test.index, idx); err != nil {
			t.Errorf(m.ErrorMessageTestCount(i+1, err))
		} else if err := a.AssertDeepEqual(test.subvector, result); err != nil {
			t.Errorf(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
