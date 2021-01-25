// +build column9_bench benchmark setbits_bench

package setbit

import (
	"math"
	"testing"

	m "github.com/LuigiAndrea/test-helper/messages"
)

func BenchmarkSetBitsTrillion(b *testing.B) {
	funcToRun := []func(uint64) int{countSetBits, countSetBitsKernighan, countSetBitsLookup, countSetBitsNibble}

	for _, f := range funcToRun {
		b.Run(m.GetFuncName(f), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(999000000000000) //999 trillions
			}
		})
	}
}

func BenchmarkSetBitsMaxUint64(b *testing.B) {
	funcToRun := []func(uint64) int{countSetBits, countSetBitsKernighan, countSetBitsLookup, countSetBitsNibble}

	for _, f := range funcToRun {
		b.Run(m.GetFuncName(f), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(math.MaxUint64)
			}
		})
	}
}
