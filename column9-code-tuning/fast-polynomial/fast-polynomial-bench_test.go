// +build all column9_bench fp_bench benchmark fastpoly_bench poly_bench

package fastpoly

import (
	"testing"

	"github.com/LuigiAndrea/test-helper/messages"
)

func BenchmarkFastPolyh(b *testing.B) {
	funcToTest := []func([]int, int) int{EvalPolynomialFast, EvalPolynomial}
	for _, f := range funcToTest {
		b.Run(messages.GetFuncName(f), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f([]int{7, 7, 5, 10, 8, 3}, 55)
			}
		})
	}
}
