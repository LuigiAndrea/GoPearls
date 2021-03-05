// +build all column4 fastsearch binarysearch fastbinary fbs

package fastsearch

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

func TestFastSearch(t *testing.T) {
	funcToTest := []func([]int, int) int{fastSearch, fastSearch2, searchFirstOccurrence}

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

	for _, f := range funcToTest {
		t.Run(m.GetFuncName(f), func(t *testing.T) {
			for i, test := range testsArray {
				if err := a.AssertDeepEqual(test.pos, f(test.input, test.v)); err != nil {
					t.Error(m.ErrorMessageTestCount(i+1, err))
				}
			}
		})
	}
}

func TestFastHashingSearch(t *testing.T) {

	for i, test := range tests {
		if err := a.AssertDeepEqual(test.pos, fastHashingSearch(test.v)); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
