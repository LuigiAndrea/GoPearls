// +build all utilities testhelper

package utilities

import (
	"math"
	"testing"
)

func TestUtilitiesTesttHelperString(t *testing.T) {
	actualValues := [][]string{{"a", "b", "c", "d", "e"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"CC", "AA", "BB"}}
	expectedValues := [][]string{{"a", "b", "c", "d", "e"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"CC", "AA", "BB"}}

	for i, actualValue := range actualValues {
		CheckArraySameValues(t, StringArrays{Expected: expectedValues[i], Actual: actualValue})
	}
}

func TestUtilitiesTesttHelperInt(t *testing.T) {
	actualValues := [][]int{{1, 32, 44322, math.MaxInt64, math.MinInt64}, {133}, {0}, {}, {-3, 43, -0}}
	expectedValues := [][]int{{1, 32, 44322, math.MaxInt64, math.MinInt64}, {133}, {0}, {}, {-3, 43, 0}}

	for i, actualValue := range actualValues {
		CheckArraySameValues(t, IntArrays{Expected: expectedValues[i], Actual: actualValue})
	}
}

func TestUtilitiesTesttHelperFloat64(t *testing.T) {
	actualValues := [][]float64{{1.0, 32.0, 44322.0, math.MaxFloat64, math.SmallestNonzeroFloat64, math.Inf(2)}, {133.0}, {0.0}, {}, {-3.0, 43.0, -0.0}}
	expectedValues := [][]float64{{1.0, 32.0, 44322.0, math.MaxFloat64, math.SmallestNonzeroFloat64, math.Inf(200)}, {133.0}, {0.0}, {}, {-3.0, 43.0, 0.0}}

	for i, actualValue := range actualValues {
		CheckArraySameValues(t, Float64Arrays{Expected: expectedValues[i], Actual: actualValue})
	}
}
