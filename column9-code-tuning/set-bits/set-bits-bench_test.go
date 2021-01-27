// +build column9_bench benchmark setbits_bench

package setbit

import (
	"math"
	"testing"

	m "github.com/LuigiAndrea/test-helper/messages"
)

var funcToRun = []func(uint64) int{countSetBits, countSetBitsKernighan, countSetBitsLookup, countSetBitsNibble}

func BenchmarkSetBitsTrillion(b *testing.B) {

	for _, f := range funcToRun {
		b.Run(m.GetFuncName(f), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(999000000000000) //999 trillions
			}
		})
	}
}

func BenchmarkSetBitsMaxUint64(b *testing.B) {

	for _, f := range funcToRun {
		b.Run(m.GetFuncName(f), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(math.MaxUint64)
			}
		})
	}
}

func BenchmarkSetBitsSequence(b *testing.B) {

	for _, f := range funcToRun {
		b.Run(m.GetFuncName(f), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				countSetBitsSequence(f)
			}
		})
	}
}
