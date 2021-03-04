// +build column4_bench benchmark fastsearch_bench fbs_bench

package fastsearch

import (
	"fmt"
	"testing"

	m "github.com/LuigiAndrea/test-helper/messages"
)

var bench_data = []int{12, 68, 434, 500}

func BenchmarkFastBinarySearch(b *testing.B) {
	funcToTest := []func([]int, int) int{fastSearch, fastSearch2, searchFirstOccurrence}

	for _, f := range funcToTest {
		for _, v := range bench_data {
			b.Run(fmt.Sprintf("%s-%v: ", m.GetFuncName(f), v), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f(data, v)
				}
			})
		}
	}
}

func BenchmarkFastHashingSearch(b *testing.B) {
	for _, v := range bench_data {
		b.Run(fmt.Sprintf("%v: ", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fastHashingSearch(v)
			}
		})
	}
}

func BenchmarkFastBinarySearchArray(b *testing.B) {
	funcToTest := []func([100]int, int) int{fastSearchArray, fastSearch2Array, searchFirstOccurrenceArray}

	for _, f := range funcToTest {
		for _, v := range bench_data {
			b.Run(fmt.Sprintf("%s-%v: ", m.GetFuncName(f), v), func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					f(dataArray, v)
				}
			})
		}
	}
}
