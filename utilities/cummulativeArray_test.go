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
		{vector: []float64{12, 0, -23, 2, -1}, expected: []float64{12, 12, -11, -9, -10}},
		{vector: []float64{-1, -4, 5}, expected: []float64{-1, -5, 0}},
		{vector: []float64{6}, expected: []float64{6}},
		{vector: []float64{}, expected: []float64{}},
		{vector: nil, expected: nil},
		{vector: []float64{6, math.Inf(1), 2}, expected: []float64{6, math.Inf(0), math.Inf(1)}},
		{vector: []float64{5, -1, math.Inf(-1), 2}, expected: []float64{5, 4, math.Inf(-1), math.Inf(-1)}},
		{vector: []float64{math.Inf(-1), math.Inf(-1)}, expected: []float64{math.Inf(-1), math.Inf(-1)}},
		{vector: []float64{math.Inf(1), math.Inf(-1)}, expected: []float64{math.Inf(1), math.NaN()}},
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
		{
			cummulativeArray: []CummulativeValue{
				{Value: 100}, {Value: -2}, {Value: -1}, {Value: 0}},
			expected: []float64{-2, -1, 0, 100}},
		{
			cummulativeArray: []CummulativeValue{
				{Value: -4}, {Value: 2.55}, {Value: 2.51}, {Value: -4.1}},
			expected: []float64{-4.1, -4, 2.51, 2.55}},
		{
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

func TestCummulativeArrayDistance(t *testing.T) {
	type testData struct {
		from, to int
		distance float64
	}

	cArray := CummulativeArray([]CummulativeValue{{Value: 100}, {Value: -2}, {Value: -1}, {Value: 0}})

	tests := []testData{
		{from: 0, to: 1, distance: 102},
		{from: 1, to: 0, distance: 102},
		{from: 1, to: 3, distance: 2},
		{from: 1, to: 4, distance: -1},
		{from: 1, to: -2, distance: -1},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(cArray.GetDistance(test.from, test.to), test.distance); err != nil {
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
