// +build all column2 rotate

package rotate

import (
	"testing"

	goth "github.com/LuigiAndrea/test-helper"
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
		t.Logf("%s", goth.GetFuncName(rotate))
		for _, test := range tests {
			actualStr := make([]string, len(test.str))
			copy(actualStr, test.str)
			rotate(actualStr, test.shift)
			if err := goth.CheckArraySameValues(goth.StringArrays{Expected: test.expectedValue, Actual: actualStr}); err != nil {
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
	if err = f(str, shiftLength); err == nil {
		t.Errorf("%s - Expected an exception of 'rotate.ShiftLeftOutOfRange'", goth.GetFuncName(f))
	} else {
		_, ok := err.(ShiftLeftOutOfRange)
		if !ok {
			t.Errorf("%s - Expected an exception %T for value '%s' and shiftLeft '%d'", goth.GetFuncName(f), err, str, shiftLength)
		} else {
			t.Logf("%s - %v", goth.GetFuncName(f), err)
		}
	}
}
