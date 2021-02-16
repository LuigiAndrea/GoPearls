// +build all column9 search sequentialSearch

package searchtuning

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	pos, v int
}

func TestSequentialSearch(t *testing.T) {

	funcToTest := []func(int) int{seqSearchOriginal, seqSearch2, seqSearch3}

	tests := []testData{
		{v: 21, pos: 6},
		{v: 2, pos: 3},
		{v: 43, pos: 5},
		{v: 8, pos: 2},
		{v: 3, pos: -1},
		{v: 5, pos: 1},
		{v: 6, pos: 4},
	}

	for _, f := range funcToTest {
		for i, test := range tests {
			if err := a.AssertDeepEqual(test.pos, f(test.v)); err != nil {
				t.Error(m.ErrorMessageTestCount(i+1, err))
			}
		}
	}
}
