// +build column9_bench benchmark charmasks_bench

package charclassification

import (
	"testing"

	m "github.com/LuigiAndrea/test-helper/messages"
)

func BenchmarkIsAlphaNum(b *testing.B) {
	funcToRun := []func(rune) bool{isAlphaNum, isAlphaNum2}

	for _, f := range funcToRun {
		b.Run(m.GetFuncName(f), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f('5')
				f('a')
				f('D')
				f('=')
			}
		})
	}
}

func BenchmarkIsDigit(b *testing.B) {
	funcToRun := []func(rune) bool{isDigit, isDigit2}

	for _, f := range funcToRun {
		b.Run(m.GetFuncName(f), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f('8')
				f('d')
			}
		})
	}
}
