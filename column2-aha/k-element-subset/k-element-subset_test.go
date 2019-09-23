// +build all column2 ksubset

package ksubset

import (
	"GoPearls/utilities"
	"testing"
)

type testData struct {
	subset        []int
	t, k          int
	expectedValue bool
}

func TestExistKSubset(t *testing.T) {

	tests := []testData{
		testData{subset: []int{11, 5, 12, 2, 18, 3, 7}, t: 11, k: 3, expectedValue: true},
		testData{subset: []int{11, 5, 12, 2, 18, 3, 7}, t: 9, k: 3, expectedValue: false},
		testData{subset: []int{11, 5, 12, 2, 18, 3, 7}, t: 11, k: 4, expectedValue: false},
		testData{subset: []int{11, 4, 7}, t: 15, k: 2, expectedValue: true},
		testData{subset: []int{11, 3, 7, 0, -14}, t: 8, k: 4, expectedValue: true},
		testData{subset: []int{11, 3, -7, 0, -14}, t: -10, k: 4, expectedValue: true},
		testData{subset: []int{11, 3, -7, 0, -14}, t: -10, k: 5, expectedValue: false},
		testData{subset: []int{11, 0, -7, 0, -14}, t: -10, k: 5, expectedValue: true},
	}

	funcToTest := []func([]int, int, int) bool{existKSubset, existKSubsetQuickSelect}

	for _, f := range funcToTest {
		for _, test := range tests {
			if r := f(test.subset, test.t, test.k); r != test.expectedValue {
				t.Errorf("\n%s: Expected value: '%t' - Actual value '%t' \nList: %#v \nt: %d, k: %d",
					utilities.GetFuncName(f), test.expectedValue, r, test.subset, test.t, test.k)
			}
		}
	}
}
