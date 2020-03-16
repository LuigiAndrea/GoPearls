// +build all column4 fastsearch binarysearch fastbinary

package fastbinary

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	array  []int
	pos, v int
}

type testDataV2 struct {
	array  [10]int
	pos, v int
}

var array = []int{1, 5, 7, 8, 9, 12, 14, 22, 390, 434}
var data = [10]int{1, 5, 7, 8, 9, 12, 14, 22, 390, 434}

func TestFastSearchVersion1(t *testing.T) {
	tests := []testData{
		testData{array: array, v: 8, pos: 3},
		testData{array: array, v: 434, pos: 9},
		testData{array: array, v: 200, pos: -1},
		testData{array: array, v: 1, pos: 0},
		testData{array: []int{}, v: 0, pos: -1},
		testData{array: []int{-5, -3, 0, 5}, v: -3, pos: 1},
		testData{array: nil, v: 0, pos: -1},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.pos, fastSearch(test.array, test.v)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestFastSearchVersion2(t *testing.T) {
	tests := []testDataV2{
		testDataV2{array: data, v: 8, pos: 3},
		testDataV2{array: data, v: 434, pos: 9},
		testDataV2{array: data, v: 200, pos: -1},
		testDataV2{array: data, v: 1, pos: 0},
		testDataV2{array: [10]int{}, v: 10, pos: -1},
		testDataV2{array: [10]int{-5, -3}, v: -3, pos: 1},
	}

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.pos, fastSearch2(test.array, test.v)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
