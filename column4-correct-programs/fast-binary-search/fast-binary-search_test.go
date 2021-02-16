// +build all column4 fastsearch binarysearch fastbinary

package fastbinary

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	input  []int
	pos, v int
}

type testDataV2 struct {
	input  [10]int
	pos, v int
}

var slice = []int{1, 5, 7, 8, 9, 12, 14, 22, 390, 434}
var array = [10]int{1, 5, 7, 8, 9, 12, 14, 22, 390, 434}

func TestFastSearchVersion1(t *testing.T) {
	tests := []testData{
		{input: slice, v: 8, pos: 3},
		{input: slice, v: 434, pos: 9},
		{input: slice, v: 390, pos: 8},
		{input: slice, v: 22, pos: 7},
		{input: slice, v: 7, pos: 2},
		{input: slice, v: 200, pos: -1},
		{input: slice, v: 1, pos: 0},
		{input: []int{}, v: 0, pos: -1},
		{input: []int{-5, -3, 0, 5}, v: -3, pos: 1},
		{input: nil, v: 0, pos: -1},
		{input: []int(nil), v: 0, pos: -1},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.pos, searchFirstOccurrence(test.input, test.v)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestFastSearch(t *testing.T) {
	funcToTest := []func([10]int, int) int{fastSearch, fastSearch2}

	tests := []testDataV2{
		{input: array, v: 8, pos: 3},
		{input: array, v: 434, pos: 9},
		{input: array, v: 390, pos: 8},
		{input: array, v: 22, pos: 7},
		{input: array, v: 7, pos: 2},
		{input: array, v: 200, pos: -1},
		{input: array, v: 1, pos: 0},
		{input: [10]int{}, v: 10, pos: -1},
		{input: [10]int{-5, -3}, v: -3, pos: 1},
	}

	for _, f := range funcToTest {
		for i, test := range tests {
			if err := a.AssertDeepEqual(test.pos, f(test.input, test.v)); err != nil {
				t.Error(m.ErrorMessageTestCount(i+1, err))
			}
		}
	}
}
