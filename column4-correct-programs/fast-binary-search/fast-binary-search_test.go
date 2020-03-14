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

func TestFastSearchVersion1(t *testing.T) {
	tests := []testData{
		testData{array: []int{1, 5, 7, 8, 9, 12, 14, 22, 434}, v: 8, pos: 3},
		testData{array: []int{1, 5, 7, 8, 9, 12, 14, 22, 434}, v: 434, pos: 8},
		testData{array: []int{1, 5, 7, 8, 9, 12, 14, 22, 434}, v: 200, pos: -1},
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
