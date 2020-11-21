// +build column9_bench benchmark search_bench

package searchtuning

import (
	"testing"
)

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
