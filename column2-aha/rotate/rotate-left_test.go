// +build all column2 rotate

package rotate

import (
	"fmt"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	str           []string
	shift         int
	expectedValue []string
}

var rotateLeftFunc = []func([]string, int) error{rotateLeftSlice, rotateLeftReverse, rotateLeftSwapRange, rotateLeftJuggling}

func TestRotateLeft(t *testing.T) {

	tests := []testData{
		{str: []string{"a", "b", "c", "d", "e"}, shift: 2, expectedValue: []string{"c", "d", "e", "a", "b"}},
		{str: []string{"c"}, shift: 1, expectedValue: []string{"c"}},
		{str: []string{""}, shift: 0, expectedValue: []string{""}},
		{str: []string{}, shift: 0, expectedValue: []string{}},
		{str: []string{"d", "e", "a"}, shift: 3, expectedValue: []string{"d", "e", "a"}},
		{str: []string{"T", "R", "E"}, shift: 0, expectedValue: []string{"T", "R", "E"}},
		{str: []string{"CC", "AA", "BB"}, shift: 2, expectedValue: []string{"BB", "CC", "AA"}},
	}

	for _, rotate := range rotateLeftFunc {
		for i, test := range tests {
			actualStr := make([]string, len(test.str))
			copy(actualStr, test.str)
			rotate(actualStr, test.shift)
			if err := a.AssertSlicesEqual(a.StringSlicesMatch{Expected: test.expectedValue, Actual: actualStr}); err != nil {
				t.Error(m.ErrorMessageTestCount(i+1, fmt.Sprintf("%s: %s", m.GetFuncName(rotate), err)))
			}
		}
	}
}

func TestShiftLeftOutOfRange(t *testing.T) {
	tests := []testData{
		{str: []string{}, shift: 2},
		{str: []string{}, shift: -8},
		{str: []string{"a", "b"}, shift: 3},
	}

	for _, rotate := range rotateLeftFunc {
		for i, test := range tests {
			if err := a.AssertDeepException(&ShiftLeftOutOfRange{shift: test.shift}, rotate(test.str, test.shift)); err != nil {
				t.Error(m.ErrorMessageTestCount(i+1, fmt.Sprintf("%s: %s", m.GetFuncName(rotate), err)))
			}
		}
	}
}
