// +build all column2 rotate

package rotate

import (
	utls "GoPearls/utilities"
	"reflect"
	"runtime"
	"strings"
	"testing"
)

func TestRotateLeftSlice(t *testing.T) {
	str := [][]string{{"a", "b", "c", "d", "e"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"CC", "AA", "BB"}}
	expectedValues := [][]string{{"c", "d", "e", "a", "b"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"BB", "CC", "AA"}}
	shiftLength := []int{2, 1, 0, 0, 3, 0, 2, 8}

	for i, v := range str {
		rotateLeftSlice(v, shiftLength[i])
		utls.CheckArraySameValues(t, utls.StringArrays{Expected: expectedValues[i], Actual: v})
	}
}

func TestRotateLeftReverse(t *testing.T) {
	str := [][]string{{"a", "b", "c", "d", "e"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"CC", "AA", "BB"}}
	expectedValues := [][]string{{"c", "d", "e", "a", "b"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"BB", "CC", "AA"}}
	shiftLength := []int{2, 1, 0, 0, 3, 0, 2}

	for i, v := range str {
		rotateLeftReverse(v, shiftLength[i])
		utls.CheckArraySameValues(t, utls.StringArrays{Expected: expectedValues[i], Actual: v})
	}
}

func TestRotateLeftSwapRange(t *testing.T) {
	str := [][]string{{"a", "b", "c", "d", "e"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"CC", "AA", "BB"}}
	expectedValues := [][]string{{"c", "d", "e", "a", "b"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"BB", "CC", "AA"}}
	shiftLength := []int{2, 1, 0, 0, 3, 0, 2}

	for i, v := range str {
		rotateLeftSwapRange(v, shiftLength[i])
		utls.CheckArraySameValues(t, utls.StringArrays{Expected: expectedValues[i], Actual: v})
	}
}

func TestRotateLeftJuggling(t *testing.T) {
	str := [][]string{{"a", "b", "c", "d", "e"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"CC", "AA", "BB"}}
	expectedValues := [][]string{{"c", "d", "e", "a", "b"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"BB", "CC", "AA"}}
	shiftLength := []int{2, 1, 0, 0, 3, 0, 2}

	for i, v := range str {
		rotateLeftJuggling(v, shiftLength[i])
		utls.CheckArraySameValues(t, utls.StringArrays{Expected: expectedValues[i], Actual: v})
	}
}

func TestShiftLeftOutOfRange(t *testing.T) {
	rotateLeftFunc := []func([]string, int) error{rotateLeftSlice, rotateLeftReverse, rotateLeftSwapRange, rotateLeftJuggling}
	for _, rotate := range rotateLeftFunc {
		shiftLeftOutOfRange(t, [][]string{{}, {}, {"a", "b"}}, []int{2, -8, 3}, rotate)
	}
}

func shiftLeftOutOfRange(t *testing.T, str [][]string, shiftLength []int, f func([]string, int) error) {
	for i, v := range str {
		var err error
		if err = f(v, shiftLength[i]); err == nil {
			t.Error("Expected an exception of 'rotate.ShiftLeftOutOfRange'")
		} else {
			_, ok := err.(ShiftLeftOutOfRange)
			if !ok {
				t.Errorf("Expected an exception %T for value '%s' and shiftLeft '%d'", err, v, shiftLength[i])
			} else {
				t.Logf("%s - %v", strings.Split(runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name(), ".")[1], err)
			}
		}
	}
}
