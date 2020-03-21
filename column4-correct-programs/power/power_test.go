// +build all column4 power

package power

import (
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	n, e, result int
}

func TestLinesBracketPoint(t *testing.T) {

	tests := []testData{
		testData{n: 2, e: 3, result: 8},
		testData{n: 2, e: 5, result: 32},
		testData{n: 2, e: 31, result: 2147483648},
		testData{n: 3, e: 3, result: 27},
		testData{n: 3, e: 2, result: 9},
		testData{n: 3, e: 0, result: 1},
		testData{n: 3, e: -2, result: -1},
		testData{n: -3, e: 2, result: 9},
		testData{n: -3, e: 3, result: -27},
		testData{n: -3, e: 1, result: -3},
		testData{n: 0, e: 5, result: 0},
	}

	for i, test := range tests {
		r := power(test.n, test.e)
		if err := a.AssertDeepEqual(test.result, r); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
