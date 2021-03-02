// +build all column9 maximumvalue maxnarray maxarray

package maximumvalue

import (
	"errors"
	"testing"

	a "github.com/LuigiAndrea/test-helper/assertions"
	m "github.com/LuigiAndrea/test-helper/messages"
)

func TestMaximumValueNArray(t *testing.T) {

	for i, test := range testsnarray {
		x, _ := NewX(test.array)
		if test.n != 0 {
			x.SetN(test.n)
		}
		if err := a.AssertDeepEqual(test.maxValue, x.Max()); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}

func TestMaximumValueNArrayExceptions(t *testing.T) {

	for i, test := range testsnarrayEx {
		x, e := NewX(test.array)
		if e == nil {
			e = x.SetN(test.n)
		}

		if err := a.AssertException(errors.New(""), e); err != nil {
			t.Error(m.ErrorMessageTestCount(i+1, err))
		}
	}
}
