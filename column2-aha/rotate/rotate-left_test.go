// +build all column2 rotate

package rotate

import (
	utls "GoPearls/utilities"
	"testing"
)

func TestRotateLeftSlice(t *testing.T) {
	str := [][]string{{"a", "b", "c", "d", "e"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"CC", "AA", "BB"}}
	expectedValues := [][]string{{"c", "d", "e", "a", "b"}, {"c"}, {""}, {}, {"d", "e", "a"}, {"T", "R", "E"}, {"BB", "CC", "AA"}}
	shiftLength := []int{2, 1, 0, 0, 3, 0, 2}

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
