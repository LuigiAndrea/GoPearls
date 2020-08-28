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
		testData{v: 21, pos: 6},
		testData{v: 2, pos: 3},
		testData{v: 43, pos: 5},
		testData{v: 3, pos: -1},
	}

	for _, f := range funcToTest {
		for i, test := range tests {
			if err := a.AssertDeepEqual(test.pos, f(test.v)); err != nil {
				t.Error(m.ErrorMessageTestCount(i+1, err))
			}
		}
	}

}

func BenchmarkSequentialSearchOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seqSearchOriginal(i)
	}
}

func BenchmarkSequentialSearch2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seqSearch2(i)
	}
}

func BenchmarkSequentialSearch3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		seqSearch3(i)
	}
}
