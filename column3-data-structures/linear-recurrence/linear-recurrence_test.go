// +build all column3 recurrence linearRecurrence

package recurrence

import (
	"GoPearls/utilities"
	"testing"
)

type testData struct {
	a              []int
	c              []int
	m              int
	expectedValues []int
}

func TestLinearRecurence(t *testing.T) {

	tests := []testData{
		testData{a: []int{0, 1, 2, 5, 6, 8, 9}, c: []int{2, 3, 4, 5, 6, 7, 9, 13},
			m: 8, expectedValues: []int{2, 3, 6, 12, 26, 47, 79, 123}},
		testData{a: []int{0, 1, 2, 5, 6, 8, 9}, c: []int{200, 3, 4, 5, 6, 7, 9, 13},
			m: 1, expectedValues: []int{200}},
		testData{a: []int{0, 1, 2, 5, 6, 8, 9}, c: []int{2, 3, 4, 5, 6, 7, 9, 13, 18, 22},
			m: 3, expectedValues: []int{2, 3, 6}},
		testData{a: nil, c: []int{12}, m: 1, expectedValues: []int{12}},
	}

	for _, test := range tests {
		if result, err := linearRecurrence(test.a, test.c, test.m); err != nil {
			t.Error(err.Error())
		} else if err = utilities.CheckArraySameValues(utilities.IntArrays{
			Expected: test.expectedValues,
			Actual:   result}); err != nil {
			t.Error(err.Error())
		}
	}
}

func TestLinearRecurenceWrongInputs(t *testing.T) {
	tests := []testData{
		testData{a: []int{1}, c: []int{}, m: 2},
		testData{a: []int{1}, c: []int{1, 3}, m: -2},
		testData{a: []int{1}, c: []int{2, 4}, m: 20},
		testData{a: []int{}, c: []int{}, m: 2},
		testData{a: nil, c: nil, m: 2},
		testData{a: []int{12}, c: nil, m: 1},
	}

	for _, test := range tests {
		if _, err := linearRecurrence(test.a, test.c, test.m); err == nil {
			t.Error("Expected inputs exception")
		}
	}
}
