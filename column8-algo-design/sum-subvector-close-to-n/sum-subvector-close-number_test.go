// +build all column8 subvectorclosenumber subvectorclosezero closezero closenumber

package sumsubvectornumber

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	vector    []float64
	subvector []float64
}

func TestSumSubVectorCloseZero(t *testing.T) {
	tests := []testData{
		testData{vector: []float64{-2.4, 5.7, -3.31, 5, 3, -2, 2}, subvector: []float64{-2, 2}},
		testData{vector: []float64{-2.4, 5.7, -3.3, 5, 3, -2, 2}, subvector: []float64{-2.4, 5.7, -3.3}},
		testData{vector: []float64{3.5, -2.4, 5.7, 3, 5, 3, -10, 2, 1}, subvector: []float64{5, 3, -10, 2}},
		testData{vector: []float64{4, -3, 3.5, 0.4, -21, 2}, subvector: []float64{0.4}},
		testData{vector: []float64{-1, -4, -0.2, -5}, subvector: []float64{-0.2}},
		testData{vector: []float64{1, -4, 5, 7, -6.5, 1.7, 0.8, -1.5}, subvector: []float64{7, -6.5}},
		testData{vector: []float64{-7, 5, 3}, subvector: []float64{-7, 5, 3}},
		testData{vector: []float64{1}, subvector: []float64{1}},
		testData{vector: []float64{}, subvector: []float64{}},
		testData{vector: nil, subvector: []float64{}},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.subvector, sumSubvectorCloseToZero(test.vector)); err != nil {
			t.Errorf(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
