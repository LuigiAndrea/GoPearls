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
	input  [100]int
	pos, v int
}

func TestFastSearch(t *testing.T) {
	funcToTest := []func([]int, int) int{fastSearch, fastSearch2, searchFirstOccurrence}

	tests := []testData{
		{input: data, v: 8, pos: 3},
		{input: data, v: 434, pos: 99},
		{input: data, v: 390, pos: 98},
		{input: data, v: 22, pos: 7},
		{input: data, v: 7, pos: 2},
		{input: data, v: 52, pos: 37},
		{input: data, v: 200, pos: -1},
		{input: data, v: 65, pos: 50},
		{input: data, v: -1, pos: 0},
	}

	for _, f := range funcToTest {
		t.Run(m.GetFuncName(f), func(t *testing.T) {
			for i, test := range tests {
				if err := a.AssertDeepEqual(test.pos, f(test.input, test.v)); err != nil {
					t.Error(m.ErrorMessageTestCount(i+1, err))
				}
			}
		})
	}
}

func TestFastSearchArray(t *testing.T) {
	funcToTest := []func([100]int, int) int{fastSearchArray, fastSearch2Array, searchFirstOccurrenceArray}

	tests := []testDataV2{
		{input: dataArray, v: 8, pos: 3},
		{input: dataArray, v: 434, pos: 99},
		{input: dataArray, v: 390, pos: 98},
		{input: dataArray, v: 22, pos: 7},
		{input: dataArray, v: 7, pos: 2},
		{input: dataArray, v: 52, pos: 37},
		{input: dataArray, v: 200, pos: -1},
		{input: dataArray, v: 65, pos: 50},
		{input: dataArray, v: -1, pos: 0},
	}

	for _, f := range funcToTest {
		t.Run(m.GetFuncName(f), func(t *testing.T) {
			for i, test := range tests {
				if err := a.AssertDeepEqual(test.pos, f(test.input, test.v)); err != nil {
					t.Error(m.ErrorMessageTestCount(i+1, err))
				}
			}
		})
	}
}
