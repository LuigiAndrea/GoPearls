// +build all column2 rotate

package rotate

import (
	"testing"

	assert "github.com/LuigiAndrea/test-helper/assertions"
	msg "github.com/LuigiAndrea/test-helper/messages"
)

type testData struct {
	str           []string
	shift         int
	expectedValue []string
}

var rotateLeftFunc = []func([]string, int) error{rotateLeftSlice, rotateLeftReverse, rotateLeftSwapRange, rotateLeftJuggling}

func TestRotateLeft(t *testing.T) {

	tests := []testData{
		testData{str: []string{"a", "b", "c", "d", "e"}, shift: 2, expectedValue: []string{"c", "d", "e", "a", "b"}},
		testData{str: []string{"c"}, shift: 1, expectedValue: []string{"c"}},
		testData{str: []string{""}, shift: 0, expectedValue: []string{""}},
		testData{str: []string{}, shift: 0, expectedValue: []string{}},
		testData{str: []string{"d", "e", "a"}, shift: 3, expectedValue: []string{"d", "e", "a"}},
		testData{str: []string{"T", "R", "E"}, shift: 0, expectedValue: []string{"T", "R", "E"}},
		testData{str: []string{"CC", "AA", "BB"}, shift: 2, expectedValue: []string{"BB", "CC", "AA"}},
	}

	for _, rotate := range rotateLeftFunc {
		t.Logf("%s", msg.GetFuncName(rotate))
		for _, test := range tests {
			actualStr := make([]string, len(test.str))
			copy(actualStr, test.str)
			rotate(actualStr, test.shift)
			if err := assert.AssertSlicesEqual(assert.StringSlicesMatch{Expected: test.expectedValue, Actual: actualStr}); err != nil {
				t.Errorf(err.Error())
			}
		}
	}
}

func TestShiftLeftOutOfRange(t *testing.T) {
	tests := []testData{
		testData{str: []string{}, shift: 2},
		testData{str: []string{}, shift: -8},
		testData{str: []string{"a", "b"}, shift: 3},
	}

	for _, rotate := range rotateLeftFunc {
		for _, test := range tests {
			shiftLeftOutOfRange(t, test.str, test.shift, rotate)
		}
	}
}

func shiftLeftOutOfRange(t *testing.T, str []string, shiftLength int, f func([]string, int) error) {
	var err error
	var ex ShiftLeftOutOfRange

	if err = f(str, shiftLength); err == nil {
		t.Errorf("%s - Expected an exception '%T' for value '%s' and shiftLeft '%d'", msg.GetFuncName(f), ex, str, shiftLength)
	} else {
		_, ok := err.(ShiftLeftOutOfRange)
		if !ok {
			t.Errorf("%s - Expected an exception '%T' for value '%s' and shiftLeft '%d'", msg.GetFuncName(f), ex, str, shiftLength)
		} else {
			t.Logf("%s - %v", msg.GetFuncName(f), err)
		}
	}
}
