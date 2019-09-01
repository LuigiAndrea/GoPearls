// +build all column2 rotate

package rotate

import (
	utls "GoPearls/utilities"
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

func TestShiftLeftOutOfRange(t *testing.T) {
	str := [][]string{{}, {}, {"a", "b"}}
	shiftLength := []int{2, -8, 3}

	for i, v := range str {
		var err error
		if err = rotateLeftReverse(v, shiftLength[i]); err == nil {
			t.Error("Expected an exception of 'rotate.ShiftLeftOutOfRange'")
		} else {
			_, ok := err.(ShiftLeftOutOfRange)
			if !ok {
				t.Errorf("Expected an exception %T for value '%s' and shiftLeft '%d'", err, v, shiftLength[i])
			}
		}
	}
}
