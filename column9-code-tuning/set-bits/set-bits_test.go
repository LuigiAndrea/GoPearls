// +build all column9 setbits

package setbit

import (
	"math"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

func TestCountSetBits(t *testing.T) {

	type testData struct {
		number uint64
		count  int
	}

	funcToTest := []func(uint64) int{countSetBits, countSetBitsKernighan, countSetBitsLookup, countSetBitsNibble}

	tests := []testData{
		testData{number: 2000, count: 6},
		testData{number: 20001234, count: 11},
		testData{number: math.MaxUint64, count: 64},
		testData{number: 0, count: 0},
	}
	for _, f := range funcToTest {
		for i, test := range tests {
			if err := a.AssertDeepEqual(test.count, f(test.number)); err != nil {
				t.Error(m.GetFuncName(f), m.ErrorMessageTestCount(i+1, err))
			}
		}
	}
}
