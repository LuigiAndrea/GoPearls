// +build all column4 terminate hailstone

package hailstone

import (
	"fmt"
	"testing"
)

func TestTerminate(t *testing.T) {
	for _, v := range []int{17, 3, 99, 77, 30, 104} {
		fmt.Printf("%3d --> %3v\n", v, hailstoneSequence(v))
	}
}
