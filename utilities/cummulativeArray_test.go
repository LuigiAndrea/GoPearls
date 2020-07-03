// +build all utilities helper cummulativeArray

package utilities

import (
	"math"
	"sort"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

func TestCummulativeArray(t *testing.T) {

	type testData struct {
		vector   []float64
		expected []float64
	}

	tests := []testData{
		testData{vector: []float64{12, 0, -23, 2, -1}, expected: []float64{12, 12, -11, -9, -10}},
		testData{vector: []float64{-1, -4, 5}, expected: []float64{-1, -5, 0}},
		testData{vector: []float64{6}, expected: []float64{6}},
		testData{vector: []float64{}, expected: []float64{}},
		testData{vector: nil, expected: nil},
		testData{vector: []float64{6, math.Inf(1), 2}, expected: []float64{6, math.Inf(0), math.Inf(1)}},
		testData{vector: []float64{5, -1, math.Inf(-1), 2}, expected: []float64{5, 4, math.Inf(-1), math.Inf(-1)}},
		testData{vector: []float64{math.Inf(-1), math.Inf(-1)}, expected: []float64{math.Inf(-1), math.Inf(-1)}},
		testData{vector: []float64{math.Inf(1), math.Inf(-1)}, expected: []float64{math.Inf(1), math.NaN()}},
	}

	for i, test := range tests {
		if err := a.AssertSlicesEqual(
			CummulativeArrayMatch{
				Expected: test.expected,
				Actual:   CalculateCummulativeArray(test.vector)}); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestCummulativeArraySort(t *testing.T) {

	type testData struct {
		cummulativeArray CummulativeArray
		expected         []float64
	}

	tests := []testData{
		testData{
			cummulativeArray: []CummulativeValue{
				{Value: 100}, {Value: -2}, {Value: -1}, {Value: 0}},
			expected: []float64{-2, -1, 0, 100}},
		testData{
			cummulativeArray: []CummulativeValue{
				{Value: -4}, {Value: 2.55}, {Value: 2.51}, {Value: -4.1}},
			expected: []float64{-4.1, -4, 2.51, 2.55}},
		testData{
			cummulativeArray: []CummulativeValue{},
			expected:         []float64{}},
	}

	for i, test := range tests {
		sort.Stable(test.cummulativeArray)
		if err := a.AssertSlicesEqual(
			CummulativeArrayMatch{
				Expected: test.expected,
				Actual:   test.cummulativeArray}); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

//Extend test-helper library for CummulativeArray test
type CummulativeArrayMatch struct {
	Expected []float64
	Actual   CummulativeArray
}

func (cum CummulativeArrayMatch) SameLength() bool { return len(cum.Expected) == len(cum.Actual) }
func (cum CummulativeArrayMatch) AreEqual(i int) bool {
	if math.IsNaN(cum.Actual[i].Value) {
		return math.IsNaN(cum.Expected[i])
	}
	return cum.Expected[i] == cum.Actual[i].Value
}
func (cum CummulativeArrayMatch) Size() int { return len(cum.Expected) }
func (cum CummulativeArrayMatch) GetError(i int) error {
	return &a.ValueError{X: cum.Expected[i], Y: cum.Actual[i].Value, Pos: i}
}
