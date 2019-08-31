// +build all column2 rotate

package rotate

import (
	utls "GoPearls/utilities"
	"testing"
)

func TestRotateLeftSlice(t *testing.T) {
	expectedValues := []string{"c", "d", "e", "a", "b"}
	str := []string{"a", "b", "c", "d", "e"}
	rotateLeftSlice(str, 2)
	utls.CheckArraySameValues(t, utls.StringArrays{Expected: expectedValues, Actual: str})
}

func TestRotateLeftReverse(t *testing.T) {
	expectedValues := []string{"c", "d", "e", "a", "b"}
	str := []string{"a", "b", "c", "d", "e"}
	rotateLeftReverse(str, 2)
	utls.CheckArraySameValues(t, utls.StringArrays{Expected: expectedValues, Actual: str})
}
