// +build all column2 ksubset

package ksubset

import (
	"fmt"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	subset        []int
	t, k          int
	expectedValue bool
}

func TestExistKSubset(t *testing.T) {

	tests := []testData{
		{subset: []int{11, 5, 12, 2, 18, 3, 7}, t: 11, k: 3, expectedValue: true},
		{subset: []int{11, 5, 12, 2, 18, 3, 7}, t: 9, k: 3, expectedValue: false},
		{subset: []int{11, 5, 12, 2, 18, 3, 7}, t: 11, k: 4, expectedValue: false},
		{subset: []int{11, 4, 7}, t: 15, k: 2, expectedValue: true},
		{subset: []int{11, 3, 7, 0, -14}, t: 8, k: 4, expectedValue: true},
		{subset: []int{11, 3, -7, 0, -14}, t: -10, k: 4, expectedValue: true},
		{subset: []int{11, 3, -7, 0, -14}, t: -10, k: 5, expectedValue: false},
		{subset: []int{11, 0, -7, 0, -14}, t: -10, k: 5, expectedValue: true},
	}

	funcToTest := []func([]int, int, int) bool{existKSubset, existKSubsetQuickSelect}

	for _, f := range funcToTest {
		for i, test := range tests {
			if err := a.AssertDeepEqual(test.expectedValue, f(test.subset, test.t, test.k)); err != nil {
				t.Error(m.ErrorMessageTestCount(i+1, fmt.Sprintf("%s: %s\nList: %#v \nt: %d, k: %d",
					m.GetFuncName(f), err, test.subset, test.t, test.k)))
			}
		}
	}
}

func TestExistKSubsetEdgeCases(t *testing.T) {
	tests := []testData{
		{subset: []int{}, t: -10, k: -2, expectedValue: false},
		{subset: nil, t: -10, k: 5, expectedValue: false},
		{subset: []int{2, 3}, t: 5, k: 3, expectedValue: false},
	}

	funcToTest := []func([]int, int, int) bool{existKSubset, existKSubsetQuickSelect}

	for _, f := range funcToTest {
		for i, test := range tests {
			if err := a.AssertDeepEqual(test.expectedValue, f(test.subset, test.t, test.k)); err != nil {
				t.Error(m.ErrorMessageTestCount(i+1, fmt.Sprintf("%s: %s\nList: %#v \nt: %d, k: %d",
					m.GetFuncName(f), err, test.subset, test.t, test.k)))
			}
		}
	}
}
