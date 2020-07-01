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

func TestSumSubVectorCloseNumberZero(t *testing.T) {
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
		if err := a.AssertDeepEqual(test.subvector, sumSubvectorCloseToNumber(test.vector, 0.0)); err != nil {
			t.Errorf(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestSumSubVectorCloseNumber(t *testing.T) {
	type testData struct {
		vector    []float64
		subvector []float64
		number    float64
	}

	tests := []testData{
		testData{vector: []float64{1, 5, 7, 8, 12}, subvector: []float64{1, 5}, number: 5.7},
		testData{vector: []float64{11, 5, 7, 12}, subvector: []float64{7, 12}, number: 18.5},
		testData{vector: []float64{6, 10, 3, 4}, subvector: []float64{3}, number: -10},
		testData{vector: []float64{-6, 10, -3, -4}, subvector: []float64{-3, -4}, number: -7.52},
		testData{vector: []float64{-6, 10, -3, -4}, subvector: []float64{10, -3}, number: 7.52},
		testData{vector: []float64{-6, 12, -4, -3, 2, 4}, subvector: []float64{-4, -3}, number: -7.0},
		testData{vector: []float64{-6, 12, -4, -3, 2, 4}, subvector: []float64{12, -4, -3, 2}, number: 7.0},
		testData{vector: []float64{-6, -12, -4, -3, -2}, subvector: []float64{-4, -3, -2}, number: -10},
		testData{vector: []float64{-1, -4, -0.2, -5}, subvector: []float64{-1}, number: -2.01},
		testData{vector: []float64{9.1, -4, -0.2, -5}, subvector: []float64{-0.2}, number: -2.01},
		testData{vector: []float64{-1, -4, -0.2, -5}, subvector: []float64{-0.2}, number: 5},
		testData{vector: []float64{2, -3.99, -2, -5}, subvector: []float64{-2}, number: -2},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.subvector, sumSubvectorCloseToNumber(test.vector, test.number)); err != nil {
			t.Errorf(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
