// +build all column4 binarysearch binarysearchdeclaration bsdeclaration declaration bsd bs

package bsdeclaration

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	array  []int
	pos, v int
}

var array = []int{1, 5, 7, 8, 9, 12, 14, 22, 390, 434}

func TestBSDeclaration(t *testing.T) {
	var bsd bsdeclaration
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
		if err := a.AssertDeepEqual(test.pos, bsd.run(test.array, test.v)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
