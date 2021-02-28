// +build column9_bench benchmark search_bench

package searchtuning

import (
	"testing"

	m "github.com/LuigiAndrea/test-helper/messages"
)

func BenchmarkSequentialSearch(b *testing.B) {
	funcToTest := []func(int) int{seqSearchOriginal, seqSearch2, seqSearch3}
	for _, f := range funcToTest {
		b.Run(m.GetFuncName(f), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(i)
			}
		})
	}
}
