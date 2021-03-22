// +build all column9 fp fastpoly poly

package fastpoly

import (
	"testing"

	"github.com/LuigiAndrea/test-helper/assertions"
	"github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	a      []int
	x      int
	result int
}

func TestEvalPoly(t *testing.T) {

	funcToTest := []func([]int, int) int{EvalPolynomial, EvalPolynomialFast}
	a := []int{7, 7, 5, 10, 8, 3}
	tests := []testData{
		{a: a, x: 3, result: 1720},
		{a: a, x: 21, result: 13903120},
		{a: a, x: 0, result: 7},
		{a: a, x: -3, result: -320},
		{a: a, x: 1, result: 40},
		{a: []int{}, x: 10, result: 0},
		{a: []int{2}, x: 10, result: 2},
	}

	for _, f := range funcToTest {
		t.Run(messages.GetFuncName(f), func(t *testing.T) {
			for i, test := range tests {
				if err := assertions.AssertDeepEqual(test.result, f(test.a, test.x)); err != nil {
					t.Error(messages.ErrorMessageTestCount(i+1, err))
				}
			}
		})
	}
}
