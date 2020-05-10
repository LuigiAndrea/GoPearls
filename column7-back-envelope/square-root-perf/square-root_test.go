// +build all column7 squareroot square root perf

package squareroot

import (
	"testing"

	"github.com/LuigiAndrea/GoPearls/utilities"
	goth "github.com/LuigiAndrea/test-helper/messages"
)

var steps int = 10000000

func TestCallSquareRoot(t *testing.T) {
	defer utilities.Elapse(goth.GetFuncName(squareRoot))()
	squareRoot(steps)
}

func TestCallSquareRootVariable(t *testing.T) {
	defer utilities.Elapse(goth.GetFuncName(squareRootVariable))()
	squareRootVariable(steps)
}
